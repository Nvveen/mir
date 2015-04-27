package sprinter_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/Nvveen/mir/containers"
	. "github.com/Nvveen/mir/sprinter"
)

// TODO Realstorage doesn't work yet?
// TODO check all Exported functions, and see if they can be made private

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
	m *mockStorage
)

type mockStorage map[string]map[string][]string

func (m *mockStorage) OpenConnection() error {
	return nil
}

func (m *mockStorage) CloseConnection() {
}

func (m *mockStorage) InsertRecord(key string, url string, collection string) (err error) {
	defer func(err *error) {
		if r := recover(); r != nil {
			*err = errors.New("could not insert into mock storage")
		}
	}(&err)
	// ignore duplicates
	for i := range (*m)[collection][key] {
		if (*m)[collection][key][i] == url {
			return nil
		}
	}
	(*m)[collection][key] = append((*m)[collection][key], url)
	return nil
}

func (m mockStorage) String() string {
	buf := new(bytes.Buffer)
	for col, _ := range m {
		fmt.Fprintf(buf, "%s:\n", col)
		for key, val := range m[col] {
			fmt.Fprintf(buf, "\t%s: ", key)
			for i := range val {
				if i < len(val)-1 {
					fmt.Fprintf(buf, "%s, ", val[i])
				} else {
					fmt.Fprintf(buf, "%s\n", val[i])
				}
			}
		}
	}
	return buf.String()
}

// We need to wrap a stringbuffer to be a ReadCloser
type strBufCloser struct {
	io.Reader
}

func (s strBufCloser) Close() error {
	return nil
}

func makeCrawler(t *testing.T) (cr *Crawler) {
	if m == nil {
		m = new(mockStorage)
	}
	*m = mockStorage{
		"urlindex":  map[string][]string{},
		"linkindex": map[string][]string{},
		"wordindex": map[string][]string{},
	}
	if c == nil {
		cr, err := NewCrawler(&containers.List{}, m)
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
	// No need for making a crawler with the correct mockStorage, as we
	// won't be using any functionality in this test.
	l := &containers.List{}
	_, err := NewCrawler(l, &mockStorage{})
	if err != nil {
		t.Fatal(err)
	}
	b := &containers.BinaryTree{}
	_, err = NewCrawler(b, &mockStorage{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetURL(t *testing.T) {
	c, err := NewCrawler(&containers.List{}, &mockStorage{})
	if err != nil {
		t.Fatal(err)
	}
	err = c.AddURL("http://www.google.com")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRetrieveHTML(t *testing.T) {
	c := makeCrawler(t)
	result, err := c.RetrieveHTML(u.String())
	if err != nil {
		t.Fatal(err)
	} else if len(result) == 0 {
		t.Fatal("no response from " + u.String())
	}
}

func TestCrawler_ExtractInfo(t *testing.T) {
	c := makeCrawler(t)
	err := c.ExtractInfo(u.String())
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

// NOTE
// If I want to split storage from sprinter, I could just move this to the
// mongo testing while importing the crawler, instead of the other way around.

func TestCrawler_RealStorage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping mongo storage testing for the crawler")
	}
	// This could be replaced with TestDB?
	db := &MongoDB{
		Host:     "127.0.0.1:40001",
		Database: "gotest",
	}
	c, err := NewCrawler(&containers.List{}, db)
	if err != nil {
		t.Fatal(err)
	}
	err = c.IndexURL("http://www.leidenuniv.nl")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCrawler_Index(t *testing.T) {
	body := `<html><head><title></title></head><body>
<a href="http://www.liacs.nl">liacs hoi</a>bla<p>wat</p></body></html>`
	c := makeCrawler(t)
	strBuf := strings.NewReader(body)
	resp := &http.Response{
		Body:    strBufCloser{strBuf},
		Request: &http.Request{URL: u},
	}
	err := c.Index(resp)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("\n%s", c)
	t.Logf("\n%s", m)
}

func TestCrawler_CheckRobots(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping because test does a HTTP request to an external site")
	}
	c := makeCrawler(t)
	u := "http://www.google.com"
	c.AddURL(u)
	if !c.CheckRobots(u) {
		t.Fatal("http://www.google.com should be indexable")
	}
	if c.CheckRobots(u + "/search") {
		t.Fatal("http://www.google.com/search should not be indexable")
	}
}

// TODO add crawler urlList pop method?
func TestCrawling(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping because actual crawling does HTTP requests")
	}
	c := makeCrawler(t)
	// Remove crawler's first url
	// Add http://www.liacs.nl
	// Start crawling
	t.Logf("%#v\n", c)
}
