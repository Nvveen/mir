package containers

// TODO getnode by index doesn't make any sense for a tree or not-array,
// so maybe retrieve it by key

type Container interface {
	AddNode(url string) (*string, error)
	GetNode(i int) (*string, error)
	RemoveNode(key string) error
	Size() int
}
