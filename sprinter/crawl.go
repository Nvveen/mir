// Package sprinter implements our fast web crawler.
package sprinter

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/Nvveen/mir/containers"
	"github.com/temoto/robotstxt-go"
	"golang.org/x/net/html"
)

// A web crawler structure that stores information on what pages
// it visits.
type Crawler struct {
	db      Storage
	urlList containers.Container
	robots  map[string]*robotstxt.RobotsData
}

var (
	ErrAddURL         = CrawlerError("unable to add URL")
	ErrNewCrawler     = CrawlerError("failed to make a Crawler object")
	ErrInvalidElement = CrawlerError("invalid element in container")
	ErrAccessDenied   = CrawlerError("not allowed to crawl this subdomain")

	SkippedFragments = []string{
		"i", "a", "the",
	}
)

const (
	RobotsSize = 2000
)

type CrawlerError string

func (e CrawlerError) Error() string {
	return "Crawler: " + string(e)
}

// Construct a new web crawler.
func NewCrawler(con containers.Container, db Storage) (c *Crawler, err error) {
	defer func() {
		if lErr := recover(); lErr != nil {
			err = ErrNewCrawler
			c = &Crawler{}
		}
	}()
	c = new(Crawler)
	c.urlList = con
	c.db = db
	c.robots = make(map[string]*robotstxt.RobotsData, RobotsSize)
	err = c.db.OpenConnection()
	if err != nil {
		return nil, err
	}
	return
}

// Add the internal base url to start crawling from.
func (c *Crawler) AddURL(URL string) (err error) {
	node, err := c.urlList.AddNode(URL)
	if err != nil {
		return err
	}
	if node == nil {
		return ErrAddURL
	}
	return
}

// Remove a URL from the container.
func (c *Crawler) RemoveURL(key string) (err error) {
	err = c.urlList.RemoveNode(key)
	return err
}

// Retrieve the HTML content of a url with index i in the Crawler.
func (c *Crawler) RetrieveHTML(key string) (result string, err error) {
	resp, err := http.Get(key)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var b []byte
	b, err = ioutil.ReadAll(resp.Body)
	result = string(b)
	return
}

// The main function that extracts various information for a url and its page.
func (c *Crawler) ExtractInfo(key string) (err error) {
	if !c.CheckRobots(key) {
		return ErrAccessDenied
	}

	resp, err := http.Get(key)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = c.IndexURL(key)
	if err != nil {
		return err
	}
	err = c.Index(resp)
	if err != nil {
		return err
	}
	return nil
}

// Split a URL into tokens and index them.
func (c *Crawler) IndexURL(u string) (err error) {
	parsed_url, err := url.Parse(u)
	if err != nil {
		return
	}
	urls, err := containers.TokenizeURL(parsed_url)
	if err != nil {
		return
	}
	for i := range urls {
		c.db.InsertRecord(urls[i], u, "urlindex")
	}
	return
}

// Index all words and checksummed links in the document and record them.
func (c *Crawler) Index(resp *http.Response) (err error) {
	z := html.NewTokenizer(resp.Body)
	insideBody := false
	key := resp.Request.URL

	for {
		tt := z.Next()
		tok := z.Token()
		switch {
		case tt == html.ErrorToken:
			// end of document
			return nil
		case tt == html.StartTagToken:
			if tok.Data == "body" {
				// We can start recording words.
				insideBody = true
			} else if tok.Data == "a" {
				err := c.indexLinks(&tok, key)
				if err != nil {
					return err
				}
			}
		case tt == html.EndTagToken:
			// Stop recording words.
			if tok.Data == "body" {
				insideBody = false
			}
		case tt == html.TextToken && insideBody:
			// We are still inside the body and we have encountered a text token, which
			// means we are finding all words inside tags.
			err = c.indexWords(&tok, resp.Request.URL.String())
			if err != nil {
				return err
			}
		}
	}
}

// Extract the ref from a link and record it after checksumming. Also add
// it to the buffer.
func (c *Crawler) indexLinks(tok *html.Token, key *url.URL) (err error) {
	for i := range tok.Attr {
		if tok.Attr[i].Key == "href" {
			val, err := url.Parse(tok.Attr[i].Val)
			if err != nil {
				continue
			}
			if !val.IsAbs() {
				val.Host = key.Host
				val.Scheme = key.Scheme
			}
			// Add to buffer.
			c.AddURL(val.String())
			sum := md5.Sum([]byte(val.String()))
			cs := hex.EncodeToString(sum[:])
			err = c.db.InsertRecord(cs, key.String(), "linkindex")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Extract all words in a body and record them.
func (c *Crawler) indexWords(tok *html.Token, u string) (err error) {
	reg := regexp.MustCompile(`\w+`)
	if !IsIgnored(tok.Data) {
		words := reg.FindAllString(tok.Data, -1)
		for i := range words {
			err = c.db.InsertRecord(words[i], u, "wordindex")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Check if the word we're checking is in the ignored fragments-list.
func IsIgnored(str string) bool {
	str = strings.ToLower(str)
	if len(str) == 0 {
		return true
	}
	for i := range SkippedFragments {
		if str == SkippedFragments[i] {
			return true
		}
	}
	return false
}

// Check for permissions with robots.txt.
func (c *Crawler) CheckRobots(u string) bool {
	// Get hostname
	URL, err := url.Parse(u)
	if err != nil {
		return false
	}
	var (
		val *robotstxt.RobotsData
		ok  bool
	)
	if val, ok = c.robots[URL.Host]; !ok {
		// Does not exist, add to map
		// But first check if map isn't full
		if len(c.robots) > RobotsSize {
			// For now discard
			// Discarding has the current added benefit that newer robots.txt
			// are retrieved again after discarding, so we don't have to check for
			// updates. Either way, might not be the best way to do this.
			c.robots = make(map[string]*robotstxt.RobotsData, RobotsSize)
		}
		r, err := http.Get("http://" + URL.Host + "/robots.txt")
		if err != nil {
			return false
		}
		rt, err := robotstxt.FromResponse(r)
		if err != nil {
			return false
		}
		c.robots[URL.Host] = rt
		val = c.robots[URL.Host]
	}
	return val.TestAgent(URL.Path, "MIR")
}
