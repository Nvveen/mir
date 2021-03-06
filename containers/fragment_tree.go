package containers

import (
	"net/url"
	"strings"
)

// TODO turn tokenization of host around (Tobias idea)

type (
	// The data structure that allows for lower storage requirements by
	// fragmenting urls and storing each fragment as a key. For this to be
	// succesfull, a large number of links need to stored (compared to a
	// simple string list.
	FragmentTree struct {
		root *fragmentNode
	}

	// A simple node structure used by the FragmentTree data structure.
	fragmentNode struct {
		label    string
		children []*fragmentNode
	}
)

type FragmentTreeError string

func (e FragmentTreeError) Error() string {
	return "Fragment Tree: " + string(e)
}

var (
	ErrTokenizer = FragmentTreeError("not a simple URL")
)

// Make a new FragmentTree object.
func NewFragmentTree() (t *FragmentTree, err error) {
	return &FragmentTree{}, nil
}

// Add a url by fragmenting it and using each fragment as a key in
// a tree.
func (f *FragmentTree) AddNode(url *url.URL) (err error) {
	words, err := TokenizeURL(url)
	if err != nil {
		return
	}
	if f.root == nil {
		f.root = new(fragmentNode)
	}
	_ = words
	var fn func(p *fragmentNode, parent *fragmentNode, i int)
	fn = func(p *fragmentNode, parent *fragmentNode, i int) {
		if i > len(words)-1 {
			return
		}
		for j := range p.children {
			if p.children[j] != nil && p.children[j].label == words[i] {
				fn(p.children[j], p, i+1)
				return
			}
		}
		// This part does not exist
		c := p
		for j := i; j < len(words); j++ {
			c.children = append(c.children, new(fragmentNode))
			c.children[len(c.children)-1].label = words[j]
			c = c.children[len(c.children)-1]
		}
	}
	fn(f.root, f.root, 0)
	return
}

func (f *FragmentTree) GetUrl(i int) (result string, err error) {
	// TODO add string of fragment parts to an internal list, use
	// fragment tree to point to each fragment, but calculate the
	// extra overhead
	return
}

// Tokenize an URL into separate alphanumeric words for indexing purposes.
func TokenizeURL(url *url.URL) (tok []string, err error) {
	if url.User != nil || len(url.Opaque) > 0 {
		err = ErrTokenizer
		return
	}
	tok = append(tok, strings.Split(url.Host, ".")...)
	// Split path
	path := url.Path
	path = strings.TrimSpace(path)
	if path != "/" && path != "" {
		if strings.HasPrefix(path, "/") {
			path = path[1:]
		}
		if strings.HasSuffix(path, "/") {
			path = path[:len(path)-1]
		}
		tok = append(tok, strings.Split(path, "/")...)
	}
	// Parse query
	values := url.Query()
	for key, val := range values {
		tok = append(tok, key)
		for i := range val {
			tok = append(tok, val[i])
		}
	}
	// Now all words in a string have been appended to the string
	return
}
