package tree

import "bytes"

// TODO maybe change package name to something more general than tree
// TODO add benchmarks

// BinaryTree is an object that uses internal structures to allow
// for sorted storage of urls.
type BinaryTree struct {
	root  *binaryNode
	nodes []string
}

// A binaryNode is the internal node type of a binary tree.
type binaryNode struct {
	label       *string
	left, right *binaryNode
}

// Add a new binary tree with an internal linked list.
func NewBinaryTree() (b *BinaryTree, err error) {
	b = new(BinaryTree)
	b.nodes = make([]string, 0, 20) // store 20 urls before having to resize
	return
}

// Add a new URL to the binary tree.
func (b *BinaryTree) AddNode(url string) (err error) {
	return b.addRecursive(&(b.root), url)
}

// A recursive adding function that adds a key string to the binary tree.
func (b *BinaryTree) addRecursive(p **binaryNode, key string) (err error) {
	if (*p) == nil {
		(*p) = new(binaryNode)
		b.nodes = append(b.nodes, key)
		(*p).label = &(b.nodes[len(b.nodes)-1])
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
