package containers

type List []string

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

func (l *List) GetNode(i int) (result *string, err error) {
	defer func() {
		if vErr := recover(); vErr != nil {
			result = nil
			err = ErrInvalidIndex
		}
	}()
	return &(*l)[i], nil
}

func (l *List) Size() int {
	return len(*l)
}
