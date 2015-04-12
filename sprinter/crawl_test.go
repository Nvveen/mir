package sprinter

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/Nvveen/mir/containers"
)

var (
	errInvalidDBElement = errors.New("Invalid Database element")

	u = &url.URL{
		Scheme:   "http",
		Opaque:   "",
		User:     nil,
		Host:     "www.leidenuniv.nl",
		Path:     "",
		RawQuery: "",
		Fragment: "",
	}
	link = "http://www.liacs.nl"

	c *Crawler
)

type mockStorage struct{}

func (m mockStorage) OpenConnection() error {
	return nil
}

func (m mockStorage) CloseConnection() {
}

func (m mockStorage) InsertRecord(key string, url string, collection string) error {
	return nil
}

// We need to wrap a stringbuffer to be a ReadCloser
type strBufCloser struct {
	io.Reader
}

func (s strBufCloser) Close() error {
	return nil
}

func makeCrawler(t *testing.T) (cr *Crawler) {
	if c == nil {
		cr, err := NewCrawler(&containers.List{}, mockStorage{})
		if err != nil {
			t.Fatal(err)
		}
		err = cr.AddURL(u.String())
		if err != nil {
			t.Fatal(err)
		}
		c = cr
	}
	return c
}

func TestNewCrawler(t *testing.T) {
	l := &containers.List{}
	_, err := NewCrawler(l, mockStorage{})
	if err != nil {
		t.Fatal(err)
	}
	b, err := containers.NewBinaryTree(&containers.List{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = NewCrawler(b, mockStorage{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetURL(t *testing.T) {
	c, err := NewCrawler(&containers.List{}, mockStorage{})
	if err != nil {
		t.Fatal(err)
	}
	err = c.AddURL("http://www.google.com")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCrawler_GetURL(t *testing.T) {
	c, err := NewCrawler(&containers.List{}, mockStorage{})
	if err != nil {
		t.Fatal(err)
	}
	err = c.AddURL("http://www.google.com")
	if err != nil {
		t.Fatal(err)
	}
	result, err := c.GetURL(0)
	if err != nil {
		t.Fatal(err)
	}
	if result != "http://www.google.com" {
		t.Fatal("invalid URL returned")
	}
}

func TestRetrieveHTML(t *testing.T) {
	c, err := NewCrawler(&containers.List{}, mockStorage{})
	if err != nil {
		t.Fatal(err)
	}
	c.AddURL("http://www.google.com")
	result, err := c.RetrieveHTML(0)
	if err != nil {
		t.Fatal(err)
	} else if len(result) == 0 {
		t.Fatal("no response from http://www.google.com")
	}
}

func TestCrawler_ExtractInfo(t *testing.T) {
	c, err := NewCrawler(&containers.List{}, mockStorage{})
	if err != nil {
		t.Fatal(err)
	}
	err = c.AddURL("http://www.leidenuniv.nl")
	if err != nil {
		t.Fatal(err)
	}
	err = c.ExtractInfo(0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCrawler_IndexURL(t *testing.T) {
	c := makeCrawler(t)
	err := c.IndexURL(u.String())
	if err != nil {
		t.Fatal(err)
	}
}

func TestCrawler_IndexLinks(t *testing.T) {
	body := `<html><head><title></title></head><body>
<a href="http://www.liacs.nl">liacs></a></body></html>`

	c := makeCrawler(t)

	// Construct an empty response
	strBuf := strings.NewReader(body)
	resp := &http.Response{
		Body:    strBufCloser{strBuf},
		Request: &http.Request{URL: u},
	}
	// Do the actual indexing
	err := c.IndexLinks(resp)
	if err != nil {
		t.Fatal(err)
	}
}
