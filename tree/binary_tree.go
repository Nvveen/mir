package tree

import (
	"bytes"
	"net/url"
)

// TODO change all url calls to just strings
// TODO add type specific errors, including to testing

// BinaryTree is an object that uses internal structures to allow
// for sorted storage of urls.
type BinaryTree struct {
	root *binaryNode
	urls []string
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
func (b *BinaryTree) AddURL(url *url.URL) (result string, err error) {
	return b.addRecursive(&(b.root), url.String())
}

// A recursive adding function that adds a key string to the binary tree.
func (b *BinaryTree) addRecursive(p **binaryNode, key string) (result string, err error) {
	if (*p) == nil {
		b.urls = append(b.urls, key)
		(*p) = new(binaryNode)
		(*p).label = &(b.urls[len(b.urls)-1])
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
