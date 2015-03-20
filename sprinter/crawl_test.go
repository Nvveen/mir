package sprinter_test

import (
	"fmt"
	"mir/sprinter"
	"testing"
)

func TestNewCrawler(t *testing.T) {
	_, err := sprinter.NewCrawler()
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewCrawlerError(t *testing.T) {
	c := sprinter.NewCrawlerError("test error")
	if c.Error() != "Crawler: test error" {
		t.Fatal("Invalid CrawlerError")
	}
}

func TestSetURL(t *testing.T) {
	c, err := sprinter.NewCrawler()
	if err != nil {
		t.Fatal(err)
	}
	err = c.AddURL("http://www.google.com")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCrawler_GetURL(t *testing.T) {
	c, err := sprinter.NewCrawler()
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
	if result.String() != "http://www.google.com" {
		t.Fatal("invalid URL returned")
	}
}

func ExampleCrawler_GetURL() {
	c, err := sprinter.NewCrawler()
	if err != nil {
		return
	}
	c.AddURL("http://www.google.com")
	result, err := c.GetURL(0)
	if err != nil {
		return
	}
	fmt.Print(result)
	// Output:
	// http://www.google.com
}

func TestRetrieveHTML(t *testing.T) {
	c, err := sprinter.NewCrawler()
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

func ExampleCrawler_RetrieveHTML(t *testing.T) {
	c, err := sprinter.NewCrawler()
	if err != nil {
		return
	}
	c.AddURL("http://www.google.com")
	result, err := c.RetrieveHTML(0)
	if err != nil {
		return
	}
	fmt.Printf("%s\n", result)
}
