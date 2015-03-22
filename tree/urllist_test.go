package tree

import (
	"errors"
	"net/url"
	"testing"
)

func makeTestList(t *testing.T) URLList {
	l := NewURLList()
	str := []*url.URL{
		&url.URL{Scheme: "http", Host: "www.google.com", Path: "/"},
		&url.URL{Scheme: "http", Host: "www.liacs.nl", Path: "/"},
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
	str := []*url.URL{
		&url.URL{Scheme: "http", Host: "www.google.com", Path: "/"},
		&url.URL{Scheme: "http", Host: "www.liacs.nl", Path: "/"},
	}
	l := makeTestList(t)
	for i := range str {
		if str[i].String() != l[i].String() {
			t.Fatal(errors.New("Adding URLS to URLList failure"))
		}
	}
}

func TestURLList_GetURL(t *testing.T) {
	l := makeTestList(t)
	url, err := l.GetURL(0)
	if err != nil {
		t.Fatal(err)
	}
	if url.String() != "http://www.google.com/" {
		t.Fatal(errors.New("Invalid URL retrieve from URLList"), ":", url)
	}
}
