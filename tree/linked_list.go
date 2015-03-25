package tree

import "errors"

// deprecated

type LinkedList struct {
	size  int
	begin *linkedListNode
	end   *linkedListNode
}

type linkedListNode struct {
	el   string
	next *linkedListNode
}

func (l *LinkedList) GetNode(i int) (result *string) {
	defer func() {
		if err := recover(); err != nil {
			panic(errors.New("invalid index in linked list"))
		}
	}()
	p := l.begin
	for j := 0; j < i; j++ {
		p = p.next
	}
	return &(p.el)
}

func (l *LinkedList) AddNode(key string) (result *string, err error) {
	c := new(linkedListNode)
	c.el = key
	result = &(c.el)
	if l.begin == nil {
		l.begin = c
		l.end = c
	} else {
		l.end.next = c
		l.end = c
	}
	l.size++
	return
}

func (l *LinkedList) Size() (s int) {
	return l.size
}
