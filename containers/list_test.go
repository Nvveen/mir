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
	// empty list
	l := &List{}
	err := l.RemoveNode("bla")
	if err == nil {
		t.Fatal("invalid attempted deletion in list")
	}
	// one element list
	l = &List{}
	l.AddNode("first")
	err = l.RemoveNode("first")
	if err != nil {
		t.Fatal(err)
	}
	if l.Size() != 0 {
		t.Fatal("invalid attempted deletion in list")
	}
	t.Logf("%s", l)
	// two element list, delete first
	l = &List{}
	l.AddNode("first")
	l.AddNode("second")
	err = l.RemoveNode("first")
	if err != nil {
		t.Fatal(err)
	}
	if l.Size() != 1 {
		t.Fatal("invalid attempted deletion in list")
	}
	t.Logf("%s", l)
	// two element list, delete second
	l = &List{}
	l.AddNode("first")
	l.AddNode("second")
	err = l.RemoveNode("second")
	if err != nil {
		t.Fatal(err)
	}
	if l.Size() != 1 {
		t.Fatal("invalid attempted deletion in list")
	}
	t.Logf("%s", l)
	// three element list
	l = &List{}
	l.AddNode("first")
	l.AddNode("second")
	l.AddNode("third")
	err = l.RemoveNode("second")
	if err != nil {
		t.Fatal(err)
	}
	if l.Size() != 2 {
		t.Fatal("invalid attempted deletion in list")
	}
	t.Logf("%s", l)
}
