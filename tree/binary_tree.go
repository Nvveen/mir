package tree

import (
	"bytes"
	"fmt"
	"net/url"
)

// TODO change all url calls to just strings
// TODO add type specific errors, including to testing

// BinaryTree is an object that uses internal structures to allow
// for sorted storage of urls.
type BinaryTree struct {
	root *binaryNode
	urls LinkedList
}

// A binaryNode is the internal node type of a binary tree.
type binaryNode struct {
	label       *string
	left, right *binaryNode
}

// Add a new binary tree with an internal linked list.
func NewBinaryTree() (b *BinaryTree, err error) {
	b = new(BinaryTree)
	return
}

// Add a new URL to the binary tree.
func (b *BinaryTree) AddURL(url *url.URL) (err error) {
	return b.addRecursive(&(b.root), url.String())
}

// A recursive adding function that adds a key string to the binary tree.
func (b *BinaryTree) addRecursive(p **binaryNode, key string) (err error) {
	if (*p) == nil {
		var c *string
		c, err = b.urls.AddNode(key)
		if err != nil {
			return
		}
		(*p) = new(binaryNode)
		(*p).label = c
		fmt.Printf("url size: %d - node: %s\n", b.urls.Size(), *((*p).label))
	} else {
		comp := bytes.Compare([]byte(key), []byte(*((*p).label)))
		if comp == -1 {
			b.addRecursive(&((*p).left), key)
		} else if comp >= 0 {
			b.addRecursive(&((*p).right), key)
		}
	}
	return
}
