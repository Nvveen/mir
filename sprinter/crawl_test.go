package sprinter

import (
	"log"
	"testing"
)

func TestNewCrawler(t *testing.T) {
	c, err := NewCrawler()
	if err != nil {
		t.Fatal(err)
	}
	if len(c.urlList) != 0 {
		t.Fatal("Invalid url list")
	}
}

func TestNewCrawlerError(t *testing.T) {
	c := NewCrawlerError("test error")
	if c.Error() != "Crawler: test error" {
		t.Fatal("Invalid CrawlerError")
	}
}

func TestSetURL(t *testing.T) {
	c, err := NewCrawler()
	if err != nil {
		t.Fatal(err)
	}
	err = c.AddURL("http://www.google.com")
	if err != nil {
		t.Fatal(err)
	}
	if c.urlList[0].String() != "http://www.google.com" {
		t.Fatal(err)
	}
}

func ExampleAddURL() {
	c, err := NewCrawler()
	if err != nil {
		log.Fatal(err)
	}
	c.AddURL("http://www.google.com")
	// Output:
}
