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

func TestBinaryTree_RemoveNode(t *testing.T) {
	b := &BinaryTree{}
	// no elements
	err := b.RemoveNode("5")
	if err == nil {
		t.Fatal("invalid attempted deletion in empty tree")
	}
	// one element
	b.AddNode("5")
	err = b.RemoveNode("5")
	if err != nil {
		t.Fatal(err)
	}
	b = &BinaryTree{}
	t.Logf("Empty tree: %s", b)
	// one element with left child
	b.AddNode("5")
	b.AddNode("4")
	err = b.RemoveNode("5")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", b)
	b = &BinaryTree{}
	// one element with right child
	b.AddNode("5")
	b.AddNode("6")
	err = b.RemoveNode("5")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", b)
	b = &BinaryTree{}
	// 	 	 5
	//    / \
	//   4   7
	//  /   / \
	// 3   6   8
	//          \
	//           9
	b.AddNode("5")
	b.AddNode("4")
	b.AddNode("3")
	b.AddNode("7")
	b.AddNode("6")
	b.AddNode("8")
	b.AddNode("9")
	err = b.RemoveNode("5")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", b)
}
