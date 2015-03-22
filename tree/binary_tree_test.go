package tree

import (
	"errors"
	"net/url"
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	_, err := NewBinaryTree()
	if err != nil {
		t.Fatal(err)
	}
}

func TestBinaryTree_AddURL(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fatal("Failed to add url")
		}
	}()
	b, err := NewBinaryTree()
	if err != nil {
		t.Fatal(err)
	}
	urls := []*url.URL{
		&url.URL{Scheme: "http", Host: "www.google.com", Path: "/"},
		&url.URL{Scheme: "http", Host: "www.liacs.nl", Path: "/"},
		&url.URL{Scheme: "http", Host: "www.bing.com", Path: "/"},
	}
	b.AddURL(urls[0])
	b.AddURL(urls[1])
	b.AddURL(urls[2])
	if *(b.root.label) != "http://www.google.com/" {
		t.Fatal(errors.New("Failed to add url"), urls[0].String())
	}
	if *(b.root.left.label) != "http://www.bing.com/" {
		t.Fatal(errors.New("Failed to add url"), urls[1].String())
	}
	if *(b.root.right.label) != "http://www.liacs.nl/" {
		t.Fatal(errors.New("Failed to add url"), urls[2].String())
	}
}
