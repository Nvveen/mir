// Package sprinter implements our fast web crawler.
package sprinter

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"golang.org/x/net/html"

	"github.com/Nvveen/mir/containers"
)

type Crawler struct {
	links                 chan string
	functionBuffer        chan int
	MaxRequests           int
	MaxConcurrentRequests int
	db                    Storage
	list                  containers.Container
}

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
	c.links = make(chan string, c.MaxConcurrentRequests)
	c.functionBuffer = make(chan int, c.MaxConcurrentRequests)
	return
}

// Start at the URI and crawl from there.
func (c *Crawler) Crawl(uri string) (err error) {
	if c.MaxRequests <= 0 || c.MaxConcurrentRequests <= 0 {
		return ErrInvalidParameters
	}
	go func() {
		c.links <- uri
	}()
	count := 0
	m := new(sync.Mutex) // a locking mechanism for the counter
L:
	for {
		if count >= c.MaxRequests {
			break L
		}
		select {
		case link := <-c.links:
			go func(mut *sync.Mutex) {
				mut.Lock()
				count++
				_, err := c.list.AddNode(link)
				mut.Unlock()
				if err != containers.ErrDuplicateElement {
					c.extractInfo(link)
				}
			}(m)
		case <-time.After(10 * time.Second):
			fmt.Printf("timed out")
			break L
		}
	}
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
	c.functionBuffer <- 1
	fmt.Printf("retrieving %s\n", link)
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := client.Get(link)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	c.index(resp)
	<-c.functionBuffer
}

// Index a response body's certain elements, including links.
func (c *Crawler) index(resp *http.Response) (err error) {
	z := html.NewTokenizer(resp.Body)
	insideBody := false
L:
	for {
		tt := z.Next()
		tok := z.Token()
		switch {
		case tt == html.ErrorToken:
			break L
		case tt == html.StartTagToken:
			if tok.Data == "body" {
				insideBody = true
			} else if tok.Data == "a" {
				for i := range tok.Attr {
					err := c.indexLinks(tok.Attr[i].Val, resp.Request.URL)
					if err != nil {
						return err
					}
				}
			}
		case tt == html.EndTagToken:
			if tok.Data == "body" {
				insideBody = false
			}
		case tt == html.TextToken && insideBody:
			//
		}
	}
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
	_, err = c.list.AddNode(u.String())
	if err != containers.ErrDuplicateElement {
		fmt.Printf("adding %s\n", u.String())
		c.links <- u.String()
		fmt.Printf("added %s\n", u.String())
	} else {
		fmt.Printf("link %s already indexed\n", u.String())
	}
	return nil
}
