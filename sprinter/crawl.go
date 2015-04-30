// Package sprinter implements our fast web crawler.
package sprinter

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/html"
)

type Crawler struct {
	links                 chan string
	MaxRequests           int
	MaxConcurrentRequests int
	db                    Storage
}

var (
	ErrInvalidParameters = errors.New("invalid parameters for crawler")
)

func NewCrawler(storage Storage) (c *Crawler, err error) {
	c = &Crawler{}
	c.MaxRequests = 10
	c.MaxConcurrentRequests = 1
	c.db = storage
	err = c.db.OpenConnection()
	if err != nil {
		return nil, err
	}
	c.links = make(chan string, c.MaxConcurrentRequests)
	return
}

func (c *Crawler) Crawl(uri string) (err error) {
	if c.MaxRequests <= 0 || c.MaxConcurrentRequests <= 0 {
		return ErrInvalidParameters
	}
	go func() {
		c.links <- uri
	}()
	count := 0
L:
	for {
		if count >= c.MaxRequests {
			break
		}
		select {
		case link, ok := <-c.links:
			count++
			if ok {
				go c.extractInfo(link)
			} else {
				fmt.Println("everything closed")
			}
		case <-time.After(5 * time.Second):
			break L
		}
	}
	return nil
}

func (c *Crawler) extractInfo(link string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("could not retrieve %s: %s\n", link, err)
		}
	}()
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
}

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
	return nil
}
