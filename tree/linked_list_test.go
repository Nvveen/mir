package tree

import "testing"

type llTestError struct {
	err string
}

var (
	errAddNode       = llTestError{err: "failed to add node"}
	errRetrieve      = llTestError{err: "failed to retrieve node"}
	errInvalidString = llTestError{err: "invalid string retrieved in linked list"}
	errInvalidSize   = llTestError{err: "invalid size for linked list"}
)

func (l llTestError) Error() string {
	return "Linked list error: " + l.err
}

func makeLinkedList(t *testing.T) *LinkedList {
	l := &LinkedList{}
	_, err := l.AddNode("http://www.google.com")
	if err != nil {
		t.Fatal(err)
	}
	_, err = l.AddNode("http://www.liacs.nl")
	if err != nil {
		t.Fatal(err)
	}
	return l
}

func TestLinkedList_AddNode(t *testing.T) {
	l := &LinkedList{}
	res, err := l.AddNode("http:/www.google.com/")
	if err != nil {
		t.Fatal(err)
	}
	if *res != l.begin.el {
		t.Fatal(errAddNode)
	}
}

func TestLinkedList_GetNode(t *testing.T) {
	l := makeLinkedList(t)
	defer func() {
		if err := recover(); err != nil {
			t.Fatal(errRetrieve)
		}
	}()
	if *(l.GetNode(0)) != "http://www.google.com" {
		t.Fatal(errInvalidString, *(l.GetNode(0)))
	}
	if *(l.GetNode(1)) != "http://www.liacs.nl" {
		t.Fatal(errInvalidString)
	}
}

func TestLinkedList_Size(t *testing.T) {
	l := makeLinkedList(t)
	if l.Size() != 2 {
		t.Fatal(errInvalidSize)
	}
}
