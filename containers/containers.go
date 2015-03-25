package containers

type Container interface {
	AddNode(url string) error
	GetNode(i int) (string, error)
}
