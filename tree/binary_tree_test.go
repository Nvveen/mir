package tree

import (
	"errors"
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
		t.Fatal(errors.New("Failed to add url"), urls[0])
	}
	if *(b.root.right.label) != "http://www.liacs.nl/" {
		t.Fatal(errors.New("Failed to add url"), urls[1])
	}
	if *(b.root.left.label) != "http://www.bing.com/" {
		t.Fatal(errors.New("Failed to add url"), urls[2])
	}
}
