package containers

import (
	"errors"
)

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

func (b *BinaryTree) AddNode(val string) (res *string, err error) {
	var n, p *binaryTreeNode
	p = insert(b.root, val, &n)
	if b.root == nil {
		b.root = p
	}
	b.size++
	return &(n.val), nil
}

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

func findmin(s **binaryTreeNode) (result **binaryTreeNode) {
	result = s
	for (*result).left != nil {
		result = &((*result).left)
	}
	return result
}

func (b *BinaryTree) Size() int {
	return b.size
}

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
