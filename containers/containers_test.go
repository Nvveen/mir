package containers

import "testing"

func TestContainer(t *testing.T) {
	var c, l Container
	l = &List{}
	c, _ = NewBinaryTree(l)
	_ = c
}
