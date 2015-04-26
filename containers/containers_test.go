package containers

import (
	"errors"
	"testing"
)

var (
	ErrUnsatisfiedInterface = errors.New("could not satisfy container interface")
)

func TestBinaryTreeContainer(t *testing.T) {
	var c Container
	c = &BinaryTree{}
	c = &List{}
	_ = c
}
