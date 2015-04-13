package containers

import "bytes"

type (
	// BinaryTree is an object that uses internal structures to allow
	// for sorted storage of urls.
	BinaryTree struct {
		root  *binaryNode
		nodes Container
	}

	// A binaryNode is the internal node type of a binary tree.
	binaryNode struct {
		label       *string
		left, right *binaryNode
	}
)

type BinaryTreeError string

func (e BinaryTreeError) Error() string {
	return "Binary Tree: " + string(e)
}

var (
	ErrInvalidIndex = BinaryTreeError("invalid index")
)

// Add a new binary tree with a backing container.
func NewBinaryTree(con Container) (b *BinaryTree, err error) {
	b = new(BinaryTree)
	b.nodes = con
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

func (b *BinaryTree) Size() int {
	return b.nodes.Size()
}
