package containers

import (
	"errors"
)

// A binary search tree
type BinaryTree struct {
	size int
	root *BinaryTreeNode
}

type BinaryTreeNode struct {
	left  *BinaryTreeNode
	val   string
	right *BinaryTreeNode
}

var (
	ErrEmptyTree       = errors.New("empty tree")
	ErrElementNotFound = errors.New("element not found")
)

func (b BinaryTreeNode) Value() string {
	return b.val
}

// Add a node to the tree.
func (b *BinaryTree) AddNode(val string) (res *string, err error) {
	var n, p *BinaryTreeNode
	p = insert(b.root, val, &n)
	if n == nil {
		return nil, ErrDuplicateElement
	} else {
		b.size++
	}
	if b.root == nil {
		b.root = p
	}
	return &(n.val), nil
}

// A recursive node-addition algorithm for a tree.
func insert(t *BinaryTreeNode, val string, n **BinaryTreeNode) *BinaryTreeNode {
	if t == nil {
		*n = &BinaryTreeNode{nil, val, nil}
		return *n
	}
	if val < t.val {
		t.left = insert(t.left, val, n)
		return t
	}
	t.right = insert(t.right, val, n)
	return t
}

// Retrieve a node by index
func (b *BinaryTree) GetNode(key string) (res *string, err error) {
	var f func(*BinaryTreeNode)
	f = func(p *BinaryTreeNode) {
		if p == nil {
			return
		}
		if key == p.val {
			res = &(p.val)
			return
		}
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
func delnode(t **BinaryTreeNode, val string, deletions *int) {
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
func findmin(s **BinaryTreeNode) (result **BinaryTreeNode) {
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
	var f func(*BinaryTreeNode, int)
	var res string
	f = func(p *BinaryTreeNode, depth int) {
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

// Iterate over each node in the tree and perform the function f.
func (b *BinaryTree) Walk(f func(Node)) {
	var g func(*BinaryTreeNode)
	g = func(p *BinaryTreeNode) {
		if p == nil {
			return
		}
		f(p)
		g(p.left)
		g(p.right)
	}
	g(b.root)
}
