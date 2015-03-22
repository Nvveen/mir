package tree

import (
	"errors"
	"testing"
)

func makeLinkedList(t *testing.T) *linkedList {
	l := &linkedList{}
	_, err := l.addNode("http://www.google.com")
	if err != nil {
		t.Fatal(err)
	}
	_, err = l.addNode("http://www.liacs.nl")
	if err != nil {
		t.Fatal(err)
	}
	return l
}

func TestlinkedList_getNode(t *testing.T) {
	l := makeLinkedList(t)
	defer func() {
		if err := recover(); err != nil {
			t.Fatal(errors.New("failed to retrieve node"))
		}
	}()
	if *(l.getNode(0)) != "http://www.google.com/" {
		t.Fatal(errors.New("invalid string retrieved from linked list"))
	}
	if *(l.getNode(1)) != "http://www.liacs.nl" {
		t.Fatal(errors.New("invalid string retrieved from linked list"))
	}
}

func TestlinkedList_size(t *testing.T) {
	l := makeLinkedList(t)
	if l.size() != 2 {
		t.Fatal(errors.New("invalid size for linked list"))
	}
}
