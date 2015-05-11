package sprinter

import (
	"github.com/Nvveen/mir/containers"
	"github.com/Nvveen/mir/storage"
	"testing"
)

func TestNewCrawler(t *testing.T) {
	c, err := NewCrawler(storage.NewMockStorage(), &containers.List{})
	if err != nil {
		t.Fatal(err)
	}
	_ = c
}

func TestCrawler_Crawl(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping because of http request")
	}
	m := storage.NewMockStorage()
	c, err := NewCrawler(m, &containers.List{})
	if err != nil {
		t.Fatal(err)
	}
	c.MaxRequests = 5
	c.MaxConcurrentRequests = 2
	c.Crawl("http://www.liacs.nl")
	t.Logf("%s\n", m)
}

func TestCrawler_CrawlSequential(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping because of http request")
	}
	m := storage.NewMockStorage()
	c, err := NewCrawler(m, &containers.List{})
	if err != nil {
		t.Fatal(err)
	}
	c.MaxRequests = 5
	c.CrawlSequential("http://www.liacs.nl")
	t.Logf("%s\n", m)
}

func TestCrawler_ConcurrentCrawl(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping because of http request")
	}
	m := storage.NewMockStorage()
	c, err := NewCrawler(m, &containers.List{})
	if err != nil {
		t.Fatal(err)
	}
	c.MaxRequests = 10
	c.MaxConcurrentRequests = 5
	c.Crawl("http://www.liacs.nl")
	t.Logf("%s\n", m)
}

func TestCrawler_InvalidLink(t *testing.T) {
	c, err := NewCrawler(storage.NewMockStorage(), &containers.List{})
	if err != nil {
		t.Fatal(err)
	}
	c.MaxRequests = 1
	c.MaxConcurrentRequests = 1
	err = c.Crawl("mailto:postmaster@localhost.com")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCrawler_robotsIgnore(t *testing.T) {
	c, _ := NewCrawler(storage.NewMockStorage(), &containers.List{})
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
