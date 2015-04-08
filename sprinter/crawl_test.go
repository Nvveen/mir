package sprinter

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/Nvveen/mir/containers"
	"gopkg.in/mgo.v2/bson"
)

// TODO move all test make-objects to a single generator function
// TODO research common testing patterns

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

	c *Crawler
)

// We need to wrap a stringbuffer to be a ReadCloser
type strBufCloser struct {
	io.Reader
}

func (s strBufCloser) Close() error {
	return nil
}

func makeCrawler(t *testing.T) (cr *Crawler) {
	if c == nil {
		cr, err := NewCrawler(&containers.List{})
		if err != nil {
			t.Fatal(err)
		}
		err = cr.AddURL(u.String())
		if err != nil {
			t.Fatal(err)
		}
		c = cr
	}
	return cr
}

func compareInput(t *testing.T, c *Crawler, collection string, checks []ReverseIndex) {
	sessionCopy := c.DB.session.Clone()
	defer sessionCopy.Close()
	col := sessionCopy.DB(c.DB.Database).C(collection)
	var results []ReverseIndex
	err := col.Find(nil).All(&results)
	if err != nil {
		t.Fatal(err)
	}
	for i := range results {
		if results[i].Key != checks[i].Key {
			t.Fatal(errInvalidDBElement)
		} else {
			for j := range results[i].URLs {
				if results[i].URLs[j] != checks[i].URLs[j] {
					t.Fatal(errInvalidDBElement)
				}
			}
		}
	}
}

func TestNewCrawler(t *testing.T) {
	l := &containers.List{}
	_, err := NewCrawler(l)
	if err != nil {
		t.Fatal(err)
	}
	b, err := containers.NewBinaryTree(&containers.List{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = NewCrawler(b)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetURL(t *testing.T) {
	c, err := NewCrawler(&containers.List{})
	if err != nil {
		t.Fatal(err)
	}
	err = c.AddURL("http://www.google.com")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCrawler_GetURL(t *testing.T) {
	c, err := NewCrawler(&containers.List{})
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
	c, err := NewCrawler(&containers.List{})
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
	c, err := NewCrawler(&containers.List{})
	if err != nil {
		t.Fatal(err)
	}
	defer cleanDB(c.DB, t)
	err = c.AddURL("http://www.leidenuniv.nl")
	if err != nil {
		t.Fatal(err)
	}
	err = c.ExtractInfo(0)
	if err != nil {
		t.Fatal(err)
	}
	// Somewhat of a testing duplicate, but doing the next check makes what
	// happens more apparent
	sessionCopy := c.DB.session.Clone()
	defer sessionCopy.Close()
	col := sessionCopy.DB(c.DB.Database).C("urlindex")
	var results []ReverseIndex
	err = col.Find(bson.M{"key": "www"}).All(&results)
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 1 {
		t.Fatalf("%s: %#v\n", errNrKeys, results)
	}
	if len(results[0].URLs) != 1 {
		t.Fatalf("%s: %#v", errNrURLs, results)
	}
}

func TestCrawler_IndexURL(t *testing.T) {
	c := makeCrawler(t)
	defer cleanDB(c.DB, t)
	err := c.IndexURL(u.String())
	if err != nil {
		t.Fatal(err)
	}
}

func TestCrawler_IndexLinks(t *testing.T) {
	body := `<html><head><title></title></head><body>
<a href="http://www.liacs.nl">liacs></a></body></html>`

	c := makeCrawler(t)
	defer cleanDB(c.DB, t)

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
	// Key isn't simple, but we can hardcode it
	// sum := md5.Sum([]byte(key))
	// cs := hex.EncodeToString(sum[:])
	cs := "77fba7c3017b4358d59ebd6dfd83bef1"
	compareInput(t, c, "linkindex", []ReverseIndex{
		ReverseIndex{Key: cs, URLs: []string{"http://www.liacs.nl"}},
	})
}
