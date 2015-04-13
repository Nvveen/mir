package containers

// a simple list of strings to contain urls.
type List []string

// Add a node to the list.
func (l *List) AddNode(url string) (result *string, err error) {
	defer func() {
		if vErr := recover(); vErr != nil {
			result = nil
			err = ErrInvalidIndex
		}
	}()
	*l = append(*l, url)
	return &(*l)[len(*l)-1], nil
}

// Retrieve a node from the list by index.
func (l *List) GetNode(i int) (result *string, err error) {
	defer func() {
		if vErr := recover(); vErr != nil {
			result = nil
			err = ErrInvalidIndex
		}
	}()
	return &(*l)[i], nil
}

// Return the size of the list.
func (l *List) Size() int {
	return len(*l)
}
