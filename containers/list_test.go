package containers

import (
	"errors"
	"testing"
)

var (
	errListInvalidLength = errors.New("Invalid length for List")
)

func makeList() List {
	l := List{
		"http://www.google.com/",
		"http://www.liacs.nl/",
		"http://www.bing.com/",
	}
	return l
}

func TestList_AddNode(t *testing.T) {
	l := makeList()
	node, err := l.AddNode("http://www.leidenuniv.nl/")
	if err != nil {
		t.Fatal(err)
	}
	if len(l) != 4 {
		t.Fatal(errListInvalidLength)
	}
	if *node != "http://www.leidenuniv.nl/" {
		t.Fatal(errInvalidElement)
	}
}

func TestList_GetNode(t *testing.T) {
	l := makeList()
	result, err := l.GetNode(0)
	if err != nil {
		t.Fatal(err)
	}
	if *result != "http://www.google.com/" {
		t.Fatal(errInvalidElement)
	}
}

func TestList_Size(t *testing.T) {
	l := makeList()
	if l.Size() != 3 {
		t.Fatal(errListInvalidLength)
	}
}
