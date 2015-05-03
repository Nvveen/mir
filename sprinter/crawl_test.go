package sprinter

import (
	"github.com/Nvveen/mir/containers"
	"testing"
)

type mockStorage map[string]map[string]string

func newMockStorage() *mockStorage {
	m := &mockStorage{
		"linkindex": map[string]string{},
		"urlindex":  map[string]string{},
		"wordindex": map[string]string{},
	}
	return m
}

func (m *mockStorage) CloseConnection() {
}

func (m *mockStorage) OpenConnection() (err error) {
	return nil
}

func (m *mockStorage) InsertRecord(key string, url string, collection string) (err error) {
	(*m)[collection][key] = url
	return nil
}

func (m *mockStorage) String() string {
	res := "{\n"
	for name, collection := range *m {
		res += name + ":\n"
		for key, val := range collection {
			res += "\t" + key + ": " + val + "\n"
		}
	}
	res += "}"
	return res
}

func TestNewCrawler(t *testing.T) {
	c, err := NewCrawler(newMockStorage(), &containers.List{})
	if err != nil {
		t.Fatal(err)
	}
	_ = c
}

func TestCrawler_Crawl(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping because of http request")
	}
	m := newMockStorage()
	c, err := NewCrawler(m, &containers.List{})
	if err != nil {
		t.Fatal(err)
	}
	c.MaxRequests = 5
	c.MaxConcurrentRequests = 1
	c.Crawl("http://www.liacs.nl")
	t.Logf("%s\n", m)
}

func TestCrawler_ConcurrentCrawl(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping because of http request")
	}
	m := newMockStorage()
	c, err := NewCrawler(m, &containers.List{})
	if err != nil {
		t.Fatal(err)
	}
	c.MaxRequests = 10
	c.MaxConcurrentRequests = 5
	c.Crawl("http://www.liacs.nl")
	t.Logf("%s\n", m)
}

func TestCrawler_robotsIgnore(t *testing.T) {
	c, _ := NewCrawler(newMockStorage(), &containers.List{})
	if c.robotsIgnore("http://www.google.com/catalogs/about") {
		t.Fatalf("invalid robots result for www.google.com/catalogs/about")
	}
	if _, ok := c.robots["www.google.com"]; !ok {
		t.Fatal("invalid robots result from www.google.com: %s", c.robots)
	}
	if !c.robotsIgnore("http://www.google.com/search") {
		t.Fatal("invalid robots result from www.google.com: %s", c.robots)
	}
}
