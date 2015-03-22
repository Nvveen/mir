package tree

import (
	"errors"
	"net/url"
)

// TODO maybe change package name to something more general than tree
// TODO add benchmarks
// TODO add URLList errors

type URLList []*url.URL

// Create a new URLList object.
func NewURLList() (list URLList) {
	return URLList{}
}

// Add a new URL object to the list.
func (u *URLList) AddURL(url *url.URL) (err error) {
	defer func() {
		if vErr := recover(); vErr != nil {
			err = errors.New("could not add URL to list")
		}
	}()
	for i := range *u {
		if *((*u)[i]) == *url {
			return
		}
	}
	*u = append(*u, url)
	return
}

// Retrieve a URL from its index.
func (u *URLList) GetURL(i int) (url *url.URL, err error) {
	defer func() {
		if dErr := recover(); dErr != nil {
			err = errors.New("could not retrieve URL from list")
		}
	}()
	url = (*u)[i]
	return
}
