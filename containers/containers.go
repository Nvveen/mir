package containers

type Container interface {
	AddNode(url string) (*string, error)
	RemoveNode(key string) error
	Size() int
}

type Node interface {
	Value() string
}
