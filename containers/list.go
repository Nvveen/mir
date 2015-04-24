package containers

import "errors"

// a simple list of strings to contain urls.

type List struct {
	root *ListNode
	back *ListNode
	size int
}

type ListNode struct {
	val  string
	next *ListNode
}

var (
	ErrListNodeNotFound = errors.New("list node not found")
)

// Add a node to the list.
func (l *List) AddNode(url string) (result *string, err error) {
	if l.root == nil {
		l.root = new(ListNode)
		l.root.val = url
		l.back = l.root
	} else {
		n := new(ListNode)
		n.val = url
		l.back.next = n
		l.back = n
	}
	l.size++
	return &(l.back.val), nil
}

// Retrieve a node from the list by index.
func (l *List) GetNode(i int) (result *string, err error) {
	if i < 0 || i >= l.size {
		return nil, ErrInvalidIndex
	}
	idx := 0
	for p := l.root; p != nil; p = p.next {
		if idx == i {
			return &(p.val), nil
		}
		idx++
	}
	return nil, ErrListNodeNotFound
}

// Return the size of the list.
func (l *List) Size() int {
	return l.size
}

// Remove the node from the list.
func (l *List) RemoveNode(i int) (err error) {
	if i < 0 || i >= l.size {
		return ErrInvalidIndex
	}
	if i == 0 {
		l.root = l.root.next
	}
	idx := 0
	for p := l.root; p != nil; p = p.next {
		if idx == i-1 {
			p.next = p.next.next
		}
		idx++
	}
	l.size--
	return nil
}

// A string representation of the list.
func (l *List) String() string {
	res := "{"
	for p := l.root; p != nil; p = p.next {
		res += p.val
		if p.next != nil {
			res += ", "
		}
	}
	res += "}"
	return res
}
