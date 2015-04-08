package sprinter

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/Nvveen/mir/containers"
	"golang.org/x/net/html"
)

// A web crawler structure that stores information on what pages
// it visits.
type Crawler struct {
	DB      *Database
	urlList containers.Container
}

var (
	ErrAddURL         = errors.New("unable to add URL")
	ErrNewCrawler     = errors.New("failed to make a Crawler object")
	ErrInvalidElement = errors.New("invalid element in container")
)

// Construct a new web crawler.
func NewCrawler(con containers.Container) (c *Crawler, err error) {
	defer func() {
		if lErr := recover(); lErr != nil {
			err = ErrNewCrawler
			c = &Crawler{}
		}
	}()
	c = new(Crawler)
	c.urlList = con
	c.DB = NewDatabase()
	// TODO Find a better solution for this
	c.DB.Database = "gotest"
	c.DB.Username = "gotestuser"
	c.DB.Password = "welcome"
	err = c.DB.OpenConnection()
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
	var b []byte
	b, err = ioutil.ReadAll(resp.Body)
	result = string(b)
	return
}

// The main function that extracts various information for a url and its page.
func (c *Crawler) ExtractInfo(i int) (err error) {
	u, err := c.GetURL(i)
	if err != nil {
		return
	}
	resp, err := http.Get(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = c.IndexURL(u)
	if err != nil {
		return
	}
	_ = resp
	return
}

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
		c.DB.InsertRecord(urls[i], u, "urlindex")
	}
	return
}

func (c *Crawler) IndexLinks(resp *http.Response) (err error) {
	key := resp.Request.URL
	sum := md5.Sum([]byte(key.String()))
	cs := hex.EncodeToString(sum[:])

	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return nil
		case html.StartTagToken, html.EndTagToken:
			tok := z.Token()
			if tok.Data == "a" {
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
						err = c.DB.InsertRecord(cs, val.String(), "linkindex")
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return
}
