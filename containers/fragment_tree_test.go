package containers_test

import (
	"errors"
	"net/url"
	"testing"

	. "github.com/Nvveen/mir/containers"
)

// TODO add examples in a new file

var (
	errInvalidTokenization = errors.New("URL Tokenize: Invalid tokenization result")
	errInvalidTree         = errors.New("Invalid fragment tree")
)

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
		t.Fatal(errInvalidTokenization)
	}
	for i := range words {
		if words[i] != compare[i] {
			t.Fatal(errInvalidTokenization)
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
	err = f.AddNode(u)
	if err != nil {
		t.Fatal(err)
	}
	u, err = url.Parse("http://www.leidenuniv.nl")
	if err != nil {
		t.Fatal(err)
	}
	err = f.AddNode(u)
	if err != nil {
		t.Fatal(err)
	}
}
