package sprinter

import (
	"net/http"
	"testing"

	"github.com/Nvveen/mir/containers"
	"gopkg.in/mgo.v2/bson"
)

// TODO move all test make-objects to a single generator function
// TODO research common testing patterns

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
	c, err := NewCrawler(&containers.List{})
	if err != nil {
		t.Fatal(err)
	}
	defer cleanDB(c.DB, t)
	err = c.AddURL("http://www.google.com")
	if err != nil {
		t.Fatal(err)
	}
	err = c.IndexURL("http://www.google.com")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCrawler_IndexLinks(t *testing.T) {
	c, err := NewCrawler(&containers.List{})
	if err != nil {
		t.Fatal(err)
	}
	defer cleanDB(c.DB, t)
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	err = c.IndexLinks(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	sessionCopy := c.DB.session.Clone()
	defer sessionCopy.Close()
	col := sessionCopy.DB(c.DB.Database).C("linkindex")
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
