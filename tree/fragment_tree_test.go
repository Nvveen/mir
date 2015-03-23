package tree

import (
	"errors"
	"net/url"
	"testing"
)

// TODO add examples in a new file

func TestTokenizeURL(t *testing.T) {
	u, err := url.Parse("http://www.google.com/?q=1&q=twee&p=wat")
	if err != nil {
		t.Fatal(err)
	}
	words, err := TokenizeURL(u)
	if err != nil {
		t.Fatal(err)
	}
	compare := []string{"www", "google", "com", "q", "1", "twee", "p", "wat"}
	if len(words) != len(compare) {
		t.Fatal("URL Tokenize: Invalid tokenization result")
	}
	for i := range words {
		if words[i] != compare[i] {
			t.Fatal("URL Tokenize: Invalid tokenization result")
		}
	}
}

func TestNewTree(t *testing.T) {
	_, err := NewFragmentTree()
	if err != nil {
		t.Fatal(err)
	}
}

func TestFragmentTree_AddURL(t *testing.T) {
	f, err := NewFragmentTree()
	if err != nil {
		t.Fatal(err)
	}
	u, err := url.Parse("http://www.google.com")
	if err != nil {
		t.Fatal(err)
	}
	err = f.AddURL(u)
	if err != nil {
		t.Fatal(err)
	}
	u, err = url.Parse("http://www.leidenuniv.nl")
	if err != nil {
		t.Fatal(err)
	}
	err = f.AddURL(u)
	if err != nil {
		t.Fatal(err)
	}
	assert := func(condition bool) {
		if !condition {
			t.Fatal(errors.New("Invalid fragment tree"))
		}
	}
	assert(f.root != nil && f.root.label == "")
	assert(
		f.root.children[0] != nil &&
			f.root.children[0].label == "www" &&
			f.root.children[0].children[0] != nil &&
			f.root.children[0].children[0].label == "google" &&
			f.root.children[0].children[1] != nil &&
			f.root.children[0].children[1].label == "leidenuniv" &&
			f.root.children[0].children[0].children[0] != nil &&
			f.root.children[0].children[0].children[0].label == "com" &&
			f.root.children[0].children[1].children[0] != nil &&
			f.root.children[0].children[1].children[0].label == "nl")
}
