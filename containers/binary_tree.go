package containers

import (
	"bytes"
	"errors"
)

// TODO add benchmarks
// TODO centralize errors/maak beter duildelijk welke errors waar horen

// BinaryTree is an object that uses internal structures to allow
// for sorted storage of urls.
type BinaryTree struct {
	root  *binaryNode
	nodes Container
}

// A binaryNode is the internal node type of a binary tree.
type binaryNode struct {
	label       *string
	left, right *binaryNode
}

var (
	ErrInvalidIndex = errors.New("Invalid index in BinaryTree")
)

// Add a new binary tree with an internal linked list.
func NewBinaryTree() (b *BinaryTree, err error) {
	b = new(BinaryTree)
	b.nodes = &List{}
	return
}

// Add a new URL to the binary tree.
func (b *BinaryTree) AddNode(url string) (result *string, err error) {
	err = b.addRecursive(&(b.root), url, &result)
	return
}

// A recursive adding function that adds a key string to the binary tree.
func (b *BinaryTree) addRecursive(p **binaryNode, key string, added **string) (err error) {
	if (*p) == nil {
		(*p) = new(binaryNode)
		node, err := b.nodes.AddNode(key)
		if err != nil {
			return err
		}
		(*p).label = node
		*added = node
	} else {
		comp := bytes.Compare([]byte(key), []byte(*((*p).label)))
		if comp == -1 {
			b.addRecursive(&((*p).left), key, added)
		} else if comp >= 0 {
			b.addRecursive(&((*p).right), key, added)
		}
	}
	return
}

func (b *BinaryTree) GetNode(i int) (result *string, err error) {
	defer func() {
		if vErr := recover(); vErr != nil {
			result = nil
			err = ErrInvalidIndex
		}
	}()
	return b.nodes.GetNode(i)
}
