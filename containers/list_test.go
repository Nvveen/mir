package containers_test

import (
	"errors"
	"testing"

	. "github.com/Nvveen/mir/containers"
)

var (
	errListInvalidLength = errors.New("Invalid length for List")
)

func TestList_AddNode(t *testing.T) {
	var l List
	res, err := l.AddNode("http://www.liacs.nl")
	if err != nil {
		t.Fatal(err)
	}
	if *res != "http://www.liacs.nl" {
		t.Fatal("invalid node added to list")
	}
}

func TestList_GetNode(t *testing.T) {
	var l List
	res, err := l.AddNode("http://www.liacs.nl")
	if err != nil {
		t.Fatal(err)
	}
	ret, err := l.GetNode(0)
	if err != nil {
		t.Fatal(err)
	}
	if *res != *ret {
		t.Fatal("inequal values in list")
	}
	_, err = l.GetNode(-1)
	if err == nil {
		t.Fatal("invalid index error expected")
	}
}

func TestList_Len(t *testing.T) {
	var l List
	_, err := l.AddNode("http://www.liacs.nl")
	if err != nil {
		t.Fatal(err)
	}
	size := l.Size()
	if size != 1 {
		t.Logf("Size: %d", size)
		t.Fatal("wrong size for list")
	}
}

func TestList_RemoveNode(t *testing.T) {
	var l List
	_, err := l.AddNode("http://www.liacs.nl")
	if err != nil {
		t.Fatal(err)
	}
	l.AddNode("http://www.liacs.nl/1")
	l.AddNode("http://www.liacs.nl/2")
	err = l.RemoveNode(0)
	if err != nil {
		t.Fatal(err)
	}
	n, err := l.GetNode(0)
	if err != nil {
		t.Fatal(err)
	}
	if *n != "http://www.liacs.nl/1" {
		t.Fatal("invalid values in list")
	}
	if l.Size() != 2 {
		t.Fatal("wrong size for list")
	}
}

func TestList_String(t *testing.T) {
	var l List
	l.AddNode("http://www.liacs.nl")
	if l.String() != "{http://www.liacs.nl}" {
		t.Fatal("invalid string format")
	}
	l.AddNode("http://www.liacs.nl")
	if l.String() != "{http://www.liacs.nl, http://www.liacs.nl}" {
		t.Fatalf("%v", l.String())
		t.Fatal("invalid string format")
	}
}
