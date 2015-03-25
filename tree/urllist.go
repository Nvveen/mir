package tree

import "errors"

// deprecated

type URLList []string

// Create a new URLList object.
func NewURLList() URLList {
	return URLList{}
}

var (
	ErrURLAdd      = errors.New("Could not add URL to urllist")
	ErrURLRetrieve = errors.New("Could not retrieve URL from list")
)

// Add a new URL object to the list.
func (u *URLList) AddURL(url string) (err error) {
	defer func() {
		if vErr := recover(); vErr != nil {
			err = ErrURLAdd
		}
	}()
	for i := range *u {
		if (*u)[i] == url {
			return
		}
	}
	*u = append(*u, url)
	return
}

// Retrieve a URL from its index.
func (u *URLList) GetURL(i int) (url string, err error) {
	defer func() {
		if dErr := recover(); dErr != nil {
			err = errors.New("could not retrieve URL from list")
		}
	}()
	url = (*u)[i]
	return
}
