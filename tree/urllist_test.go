package tree

import (
	"errors"
	"testing"
)

var (
	errURLAdd = errors.New("Adding URLs to URLList failed")
)

func makeTestList(t *testing.T) URLList {
	l := NewURLList()
	str := []string{
		"http://www.google.com/",
		"http://www.liacs.nl/",
	}
	err := l.AddURL(str[0])
	if err != nil {
		t.Fatal(err)
	}
	err = l.AddURL(str[1])
	if err != nil {
		t.Fatal(err)
	}
	return l
}

func TestURLList_AddURL(t *testing.T) {
	str := []string{
		"http://www.google.com/",
		"http://www.liacs.nl/",
	}
	l := makeTestList(t)
	for i := range str {
		if str[i] != l[i] {
			t.Fatal(errURLAdd)
		}
	}
}

func TestURLList_GetURL(t *testing.T) {
	l := makeTestList(t)
	url, err := l.GetURL(0)
	if err != nil {
		t.Fatal(err)
	}
	if url != "http://www.google.com/" {
		t.Fatal(errURLAdd)
	}
}
