package tree

import (
	"errors"
	"testing"
)

var (
	errTreeError = errors.New("Invalid tree error returned")
	errURLFail   = errors.New("Failed to add URL")
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
			t.Fatal(errURLFail)
		}
	}()
	b, err := NewBinaryTree()
	if err != nil {
		t.Fatal(err)
	}
	urls := []string{
		"http://www.google.com/",
		"http://www.liacs.nl/",
		"http://www.bing.com/",
	}
	err = b.AddNode(urls[0])
	if err != nil {
		t.Fatal(err)
	}
	err = b.AddNode(urls[1])
	if err != nil {
		t.Fatal(err)
	}
	err = b.AddNode(urls[2])
	if err != nil {
		t.Fatal(err)
	}

	if *(b.root.label) != "http://www.google.com/" {
		t.Fatal(errURLFail)
	}
	if *(b.root.right.label) != "http://www.liacs.nl/" {
		t.Fatal(errURLFail)
	}
	if *(b.root.left.label) != "http://www.bing.com/" {
		t.Fatal(errURLFail)
	}
}
