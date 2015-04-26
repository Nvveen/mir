package containers

import "errors"

// a simple list of strings to contain urls.
// TODO fucking fix errors nou eens een keer

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
	ErrInvalidIndex     = errors.New("invalid index")
	ErrListNodeNotFound = errors.New("list node not found")
	ErrEmptyList        = errors.New("empty list")
)

// Add a node to the list.
func (l *List) AddNode(key string) (result *string, err error) {
	n := new(ListNode)
	n.val = key
	if l.back != nil {
		l.back.next = n
	}
	l.back = n
	if l.root == nil {
		l.root = n
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
func (l *List) RemoveNode(key string) (err error) {
	p := l.root
	if p == nil {
		return ErrEmptyList
	} else if p.val == key {
		l.root = l.root.next
		l.size--
		return nil
	}
	pn := l.root.next
	for pn != nil {
		if pn.val == key {
			p.next = pn.next
			if pn == l.back {
				l.back = p
			}
			l.size--
			return nil
		}
		p = p.next
		pn = p.next.next
	}
	return ErrElementNotFound
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
