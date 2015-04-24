package containers

import "fmt"

// TODO add comments, proper error returns

type BinaryTree struct {
	size int
	root *binaryTreeNode
}

type binaryTreeNode struct {
	left  *binaryTreeNode
	val   string
	right *binaryTreeNode
}

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
	return res, nil
}

func (b *BinaryTree) RemoveNode(val string) (err error) {
	if b.root == nil {
		return nil
	}
	if b.root.left == nil && b.root.right == nil {
		b.root = nil
		return
	}
	prev := b.size
	if b.root.left != nil {
		delnode(b.root.left, &(b.root), val)
		if prev != b.size {
			return nil
		}
	}
	if b.root.right != nil {
		delnode(b.root.right, &(b.root), val)
	}
	fmt.Printf("%#v\n", b.root)
	return nil
}

func delnode(t *binaryTreeNode, parent **binaryTreeNode, val string) {
	if t == nil {
		return
	}
	if val == t.val && t.left == nil && t.right == nil {
		if (*parent).left == t {
			(*parent).left = nil
		} else if (*parent).right == t {
			(*parent).right = nil
		}
	} else if val == t.val && t.left != nil && t.right == nil {
		if (*parent).left == t {
			(*parent).left = t.left
		} else if (*parent).right == t {
			(*parent).right = t.left
		}
	} else if val == t.val && t.left == nil && t.right != nil {
		if (*parent).left == t {
			(*parent).left = t.right
		} else if (*parent).right == t {
			(*parent).right = t.right
		}
	} else if val == t.val && t.left != nil && t.right != nil {
		t.val = t.left.val
		delnode(t.left, &t, val)
	}
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
