// Package sprinter implements our fast web crawler.
package sprinter

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"

	"github.com/Nvveen/mir/containers"
)

// TODO readd robots.txt
// TODO memory can grow indefinitely

type Crawler struct {
	links                 chan string
	functionBuffer        chan bool
	MaxRequests           int // The max number of requests that can be handled in total.
	MaxConcurrentRequests int // The max number of requests that can be handled concurrently.
	db                    Storage
	list                  containers.Container
}

var (
	ErrInvalidParameters = errors.New("invalid parameters for crawler")
	IgnoredWords         = []string{"the"}
)

// Create a new Crawler object with the specified Storage and link buffer.
func NewCrawler(storage Storage, buffer containers.Container) (c *Crawler, err error) {
	c = &Crawler{}
	c.MaxRequests = 10
	c.MaxConcurrentRequests = 1
	c.db = storage
	err = c.db.OpenConnection()
	if err != nil {
		return nil, err
	}
	c.list = buffer
	return
}

// Start at the URI and crawl from there.
func (c *Crawler) Crawl(uri string) (err error) {
	if c.MaxRequests <= 0 || c.MaxConcurrentRequests <= 0 {
		return ErrInvalidParameters
	}
	c.links = make(chan string, c.MaxConcurrentRequests)
	c.functionBuffer = make(chan bool, c.MaxConcurrentRequests)
	go func() {
		c.links <- uri
	}()
	var wg sync.WaitGroup
	wg.Add(c.MaxRequests)
	for count := 0; count < c.MaxRequests; count++ {
		link := <-c.links
		go func() {
			defer wg.Done()
			c.extractInfo(link)
		}()
	}
	wg.Wait()
	return nil
}

// A concurrent method to retrieve a HTTP object's body, and extract the
// necessary information, such as links and more.
func (c *Crawler) extractInfo(link string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("could not retrieve %s: %s\n", link, err)
		}
	}()
	c.functionBuffer <- true
	fmt.Printf("retrieving %s\n", link)
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := client.Get(link)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	err = c.indexURL(link)
	if err != nil {
		panic(err)
	}
	c.indexContent(resp)
	<-c.functionBuffer
}

// Index a response body's certain elements, including links.
func (c *Crawler) indexContent(resp *http.Response) (err error) {
	defer func(rErr *error) {
		if err := recover(); err != nil {
			*rErr = err.(error)
		}
	}(&err)
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	var f, g func(*html.Node)
	insideBody := false
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for i := range n.Attr {
				if n.Attr[i].Key == "href" {
					err := c.indexLinks(n.Attr[i].Val, resp.Request.URL)
					if err != nil {
						panic(err)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	g = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "body" {
			insideBody = true
		}
		if insideBody && n.Type == html.TextNode {
			err := c.indexWords(n.Data, resp.Request.URL)
			if err != nil {
				panic(err)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if insideBody && c.Type == html.ElementNode && c.Data == "script" {
				n.RemoveChild(c)
				continue
			}
			g(c)
		}
	}
	f(doc)
	g(doc)
	return nil
}

// Index a link into the storage object.
func (c *Crawler) indexLinks(uri string, key *url.URL) (err error) {
	u, err := url.Parse(uri)
	if err != nil {
		return err
	}
	if !u.IsAbs() {
		u.Host = key.Host
		u.Scheme = key.Scheme
	}
	sum := md5.Sum([]byte(u.String()))
	cs := hex.EncodeToString(sum[:])
	err = c.db.InsertRecord(cs, key.String(), "linkindex")
	if err != nil {
		return err
	}
	// We have indexed the link, now add it to the sprinter to be crawled,
	// after we have determined it is not a duplicate (by adding it to the list,
	// which should have duplicate detection.
	// Take note that this next bit may block, but it doesn't matter as the
	// indexing is still valid.
	_, err = c.list.AddNode(u.String())
	if err != containers.ErrDuplicateElement {
		go func() {
			c.links <- u.String()
		}()
	}
	return nil
}

// Index a URL by splitting it into components and indexing each one.
func (c *Crawler) indexURL(uri string) (err error) {
	parsed_url, err := url.Parse(uri)
	if err != nil {
		return err
	}
	urls, err := containers.TokenizeURL(parsed_url)
	if err != nil {
		return err
	}
	for i := range urls {
		c.db.InsertRecord(urls[i], uri, "urlindex")
	}
	return nil
}

// Take the body of a HTTP Get request, parse it for singular words, normalize
// them and insert them into the database.
func (c *Crawler) indexWords(data string, uri *url.URL) (err error) {
	reg := regexp.MustCompile(`\w+`)
	words := reg.FindAllString(data, -1)
L:
	for i := range words {
		// don't save single letter words
		if len(words[i]) == 1 {
			continue
		}
		// ignore capitalization
		w := strings.ToLower(words[i])
		// filter for ignored words
		for j := range IgnoredWords {
			if w == IgnoredWords[j] {
				continue L
			}
		}
		err := c.db.InsertRecord(w, uri.String(), "wordindex")
		if err != nil {
			return err
		}
	}
	return nil
}
