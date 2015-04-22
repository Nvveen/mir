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
	"golang.org/x/net/html"
)

// TODO Merge indexing links and words, because it won't work concurrently

// A web crawler structure that stores information on what pages
// it visits.
type Crawler struct {
	db      Storage
	urlList containers.Container
}

var (
	ErrAddURL         = CrawlerError("unable to add URL")
	ErrNewCrawler     = CrawlerError("failed to make a Crawler object")
	ErrInvalidElement = CrawlerError("invalid element in container")

	SkippedFragments = []string{
		"i", "a", "the",
	}
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

// Retrieve a url from the crawler's list by index.
func (c *Crawler) GetURL(i int) (ret string, err error) {
	defer func() {
		if err := recover(); err != nil {
			err = ErrInvalidElement
			c = nil
		}
	}()
	val, err := c.urlList.GetNode(i)
	if err != nil {
		return "", err
	}
	if val == nil {
		return "", ErrInvalidElement
	}
	return *val, err
}

// Retrieve the HTML content of a url with index i in the Crawler.
func (c *Crawler) RetrieveHTML(i int) (result string, err error) {
	url, err := c.GetURL(i)
	if err != nil {
		return
	}
	resp, err := http.Get(url)
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
func (c *Crawler) ExtractInfo(i int) (err error) {
	u, err := c.GetURL(i)
	if err != nil {
		return err
	}
	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = c.IndexURL(u)
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

// Extract the ref from a link and record it after checksumming.
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
