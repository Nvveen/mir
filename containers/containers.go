package containers

type Container interface {
	AddNode(url string) (*string, error)
	GetNode(i int) (*string, error)
	Size() int
}