package tree

import "errors"

type linkedList struct {
	begin *linkedListNode
	end   *linkedListNode
}

type linkedListNode struct {
	el   string
	next *linkedListNode
}

func (l *linkedList) getNode(i int) (result *string) {
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

func (l *linkedList) addNode(key string) (result *string, err error) {
	c := new(linkedListNode)
	c.el = key
	result = &(c.el)
	if l.begin == nil {
		l.begin = c
	}
	l.end = c
	return
}

func (l *linkedList) size() (s int) {
	var p *linkedListNode
	s = -1
	for p = l.begin; p != nil; p = p.next {
		s++
	}
	return s
}
