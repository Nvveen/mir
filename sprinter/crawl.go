package sprinter

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

// TODO change urllist to our custom type

// A web crawler structure that stores information on what pages
// it visits.
type Crawler struct {
	urlList []*url.URL
}

// A web crawler error type.
type CrawlerError struct {
	err string
}

// Construct a new crawler error.
func NewCrawlerError(err string) (c CrawlerError) {
	c.err = err
	return
}

// Error reports an error.
func (c CrawlerError) Error() string {
	return "Crawler: " + c.err
}

// Construct a new web crawler.
func NewCrawler() (c Crawler, err error) {
	defer func() {
		if lErr := recover(); lErr != nil {
			err = NewCrawlerError("Failed to make a Crawler object")
			c = Crawler{}
		}
	}()
	return
}

// Add the internal base url to start crawling from.
func (c *Crawler) AddURL(URL string) (err error) {
	var parsedUrl *url.URL
	parsedUrl, err = url.Parse(URL)
	if err != nil {
		return
	}
	c.urlList = append(c.urlList, parsedUrl)
	return
}

// Retrieve a url from the crawler's list by index.
func (c *Crawler) GetURL(i int) (ret *url.URL, err error) {
	defer func() {
		if err := recover(); err != nil {
			err = NewCrawlerError("Invalid element in url list")
			c = nil
		}
	}()
	ret = c.urlList[i]
	if ret == nil {
		err = CrawlerError{err: "Invalid element in url list"}
	}
	return
}

// Retrieve the HTML content of a url with index i in the Crawler.
func (c *Crawler) RetrieveHTML(i int) (result string, err error) {
	url, err := c.GetURL(i)
	if err != nil {
		return
	}
	resp, err := http.Get(url.String())
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
	url, err := c.GetURL(i)
	if err != nil {
		return
	}
	resp, err := http.Get(url.String())
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return
	}
	var walker func(*html.Node) error
	walker = func(n *html.Node) (err error) {
		// Extract links
		if n.Type == html.ElementNode && n.Data == "a" {
			for i := range n.Attr {
				if n.Attr[i].Key == "href" {
					err = c.AddURL(n.Attr[i].Val)
					if err != nil {
						return err
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walker(c)
		}
		return
	}
	walker(doc)
	return
}
