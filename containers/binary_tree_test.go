package containers_test

import (
	"errors"
	"testing"

	. "github.com/Nvveen/mir/containers"
)

var (
	errTreeError      = errors.New("Invalid tree")
	errURLFail        = errors.New("Failed to add URL")
	errInvalidError   = errors.New("Invalid error returned")
	errInvalidElement = errors.New("Invalid element from tree")
)

type mockList struct {
}

func (m *mockList) AddNode(url string) (*string, error) {
	return nil, nil
}

func (m *mockList) GetNode(i int) (*string, error) {
	return nil, nil
}

func (m *mockList) RemoveNode(i int) error {
	return nil
}

func (m *mockList) Size() int {
	return 0
}

func TestNewBinaryTree(t *testing.T) {
	_, err := NewBinaryTree(&mockList{})
	if err != nil {
		t.Fatal(err)
	}
}

func makeBinaryTree(t *testing.T) *BinaryTree {
	b, err := NewBinaryTree(&mockList{})
	if err != nil {
		t.Fatal(err)
	}
	urls := []string{
		"http://www.google.com/",
		"http://www.liacs.nl/",
		"http://www.bing.com/",
	}
	node, err := b.AddNode(urls[0])
	if err != nil {
		t.Fatal(err)
	}
	if node == nil || *node != urls[0] {
		t.Fatal(errInvalidElement)
	}
	node, err = b.AddNode(urls[1])
	if err != nil || node == nil || *node != urls[1] {
		t.Fatal(err)
	}
	node, err = b.AddNode(urls[2])
	if err != nil || node == nil || *node != urls[2] {
		t.Fatal(err)
	}
	return b
}

func TestBinaryTree_AddURL(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fatal(errURLFail)
		}
	}()
	makeBinaryTree(t)
}

func TestBinaryTree_GetNode(t *testing.T) {
	b := makeBinaryTree(t)
	result, err := b.GetNode(0)
	if err != nil {
		t.Fatal(err)
	}
	if *result != "http://www.google.com/" {
		t.Fatal(errInvalidElement)
	}
	result, err = b.GetNode(1)
	if err != nil {
		t.Fatal(err)
	}
	if *result != "http://www.liacs.nl/" {
		t.Fatal(err)
	}
	result, err = b.GetNode(3) // does not exist
	if err != ErrInvalidIndex {
		t.Fatal(errInvalidError)
	}
}
