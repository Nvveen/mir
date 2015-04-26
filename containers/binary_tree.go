package containers

import (
	"errors"
)

// TODO change GetNode to retrieve by key

// A binary search tree
type BinaryTree struct {
	size int
	root *binaryTreeNode
}

type binaryTreeNode struct {
	left  *binaryTreeNode
	val   string
	right *binaryTreeNode
}

var (
	ErrEmptyTree       = errors.New("empty tree")
	ErrElementNotFound = errors.New("element not found")
)

// Add a node to the tree.
func (b *BinaryTree) AddNode(val string) (res *string, err error) {
	var n, p *binaryTreeNode
	p = insert(b.root, val, &n)
	if b.root == nil {
		b.root = p
	}
	b.size++
	return &(n.val), nil
}

// A recursive node-addition algorithm for a tree.
func insert(t *binaryTreeNode, val string, n **binaryTreeNode) *binaryTreeNode {
	if t == nil {
		*n = &binaryTreeNode{nil, val, nil}
		return *n
	}
	if val < t.val {
		t.left = insert(t.left, val, n)
		return t
	}
	t.right = insert(t.right, val, n)
	return t
}

// Retrieve a node by index TODO will be by key
func (b *BinaryTree) GetNode(i int) (res *string, err error) {
	var f func(*binaryTreeNode)
	idx := 0
	f = func(p *binaryTreeNode) {
		if p == nil {
			return
		}
		if idx == i {
			res = &(p.val)
			return
		}
		idx++
		f(p.left)
		f(p.right)
	}
	f(b.root)
	if res == nil {
		return nil, ErrElementNotFound
	}
	return res, nil
}

// Remove a node by key from the tree. (oh yeah, it rhymes, deal with it).
func (b *BinaryTree) RemoveNode(key string) (err error) {
	if b.root == nil {
		return ErrEmptyTree
	}
	deletions := 0
	delnode(&(b.root), key, &deletions)
	if deletions == 0 {
		return ErrElementNotFound
	}
	return nil
}

// Recursive deletion in a binary tree, see wikipedia for the algorithm.
func delnode(t **binaryTreeNode, val string, deletions *int) {
	if t == nil {
		return
	}
	if val < (*t).val {
		delnode(&((*t).left), val, deletions)
	} else if val > (*t).val {
		delnode(&((*t).right), val, deletions)
	} else {
		(*deletions)++
		// delete key
		if (*t).left != nil && (*t).right != nil {
			successor := findmin(&((*t).right))
			(*t).val = (*successor).val
			delnode(successor, (*successor).val, deletions)
		} else if (*t).left != nil {
			*t = (*t).left
		} else if (*t).right != nil {
			*t = (*t).right
		} else {
			*t = nil
		}
	}
}

// Find a minimal successor node in the tree.
func findmin(s **binaryTreeNode) (result **binaryTreeNode) {
	result = s
	for (*result).left != nil {
		result = &((*result).left)
	}
	return result
}

// Return the size of the tree (nr. of nodes).
func (b *BinaryTree) Size() int {
	return b.size
}

// A string representation of the binary tree.
func (b *BinaryTree) String() string {
	var f func(*binaryTreeNode, int)
	var res string
	f = func(p *binaryTreeNode, depth int) {
		if p == nil {
			return
		}
		for d := depth; d > 0; d-- {
			res += "-"
		}
		res += p.val + "\n"
		f(p.left, depth+1)
		f(p.right, depth+1)
	}
	f(b.root, 0)
	return res
}
