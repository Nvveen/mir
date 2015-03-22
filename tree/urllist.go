package tree

import "net/url"

// TODO maybe change package name to something more general than tree
// TODO add benchmarks

type URLList []*url.URL

// Create a new URLList object.
func NewURLList() (list *URLList) {
	return &URLList{}
}

// Add a new URL object to the list.
func (u *URLList) AddURL(url *url.URL) (err error) {
	defer func() {
		err = recover()
	}()
	*u = append(*u, url)
}
