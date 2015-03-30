package sprinter

import (
	"fmt"
	"github.com/Nvveen/mir/containers"
	"testing"
)

// TODO move all test make-objects to a single generator function

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

func ExampleCrawler_GetURL() {
	c, err := NewCrawler(&containers.List{})
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

func ExampleCrawler_RetrieveHTML(t *testing.T) {
	c, err := NewCrawler(&containers.List{})
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

func TestCrawler_ExtractInfo(t *testing.T) {
	c, err := NewCrawler(&containers.List{})
	if err != nil {
		t.Fatal(err)
	}
	c.AddURL("http://www.leidenuniv.nl")
	err = c.ExtractInfo(0)
	if err != nil {
		t.Fatal(err)
	}
}
