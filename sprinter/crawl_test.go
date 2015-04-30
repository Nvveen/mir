package sprinter_test

import "testing"
import (
	"github.com/Nvveen/mir/containers"
	. "github.com/Nvveen/mir/sprinter"
)

type mockStorage map[string]map[string]string

func newMockStorage() *mockStorage {
	m := &mockStorage{
		"linkindex": map[string]string{},
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
	c.MaxRequests = 100
	c.MaxConcurrentRequests = 1
	c.Crawl("http://www.liacs.nl")
	t.Logf("%#v\n", m)
}
