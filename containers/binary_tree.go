package containers

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
