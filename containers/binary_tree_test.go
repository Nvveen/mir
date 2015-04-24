package containers_test

import "testing"
import . "github.com/Nvveen/mir/containers"

func TestBinaryTree_AddNode(t *testing.T) {
	b := &BinaryTree{}
	res, err := b.AddNode("http://www.liacs.nl")
	if err != nil {
		t.Fatal(err)
	}
	if *res != "http://www.liacs.nl" {
		t.Fatal("invalid element added in binary tree")
	}
	if b.Size() != 1 {
		t.Fatal("invalid size for binary tree")
	}
	res, err = b.AddNode("http://www.leidenuniv.nl")
	if err != nil {
		t.Fatal(err)
	}
	if *res != "http://www.leidenuniv.nl" {
		t.Logf("%s", *res)
		t.Fatal("invalid element added in binary tree")
	}
	if b.Size() != 2 {
		t.Fatal("invalid size for binary tree")
	}
}

func TestBinaryTree_String(t *testing.T) {
	b := &BinaryTree{}
	b.AddNode("http://www.leidenuniv.nl")
	b.AddNode("http://www.liacs.nl")
	b.AddNode("http://www.alpha.nl")
	b.AddNode("http://www.beta.nl")
	t.Logf("%s", b)
}

func TestBinaryTree_GetNode(t *testing.T) {
	b := &BinaryTree{}
	b.AddNode("http://www.leidenuniv.nl")
	res, err := b.GetNode(0)
	if err != nil {
		t.Fatal(err)
	}
	if *res != "http://www.leidenuniv.nl" {
		t.Fatal("invalid element retrieved from binary tree")
	}
	b.AddNode("http://www.liacs.nl")
	res, err = b.GetNode(1)
	if err != nil {
		t.Fatal(err)
	}
	if *res != "http://www.liacs.nl" {
		t.Fatal("invalid element retrieved from binary tree")
	}
}
