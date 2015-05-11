package sprinter

import (
	"testing"

	"github.com/Nvveen/mir/containers"
)

// Sequential <-> Concurrent
// Static <-> Actual MongoDB backend
// List <-> Binary Search Tree
// Remember to determine the parallelism.

func BenchmarkCrawl_SequentialStaticList(b *testing.B) {
	c, err := NewCrawler(newMockStorage(), &containers.List{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = 2000
	c.MaxConcurrentRequests = 1
	b.ResetTimer()
	c.CrawlSequential("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent10StaticList(b *testing.B) {
	c, err := NewCrawler(newMockStorage(), &containers.List{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = 2000
	c.MaxConcurrentRequests = 10
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent100StaticList(b *testing.B) {
	c, err := NewCrawler(newMockStorage(), &containers.List{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = 2000
	c.MaxConcurrentRequests = 100
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent1000StaticList(b *testing.B) {
	c, err := NewCrawler(newMockStorage(), &containers.List{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = 2000
	c.MaxConcurrentRequests = 1000
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent1000MongoList(b *testing.B) {
	c, err := NewCrawler(TestDB, &containers.List{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = 2000
	c.MaxConcurrentRequests = 1000
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent1000MongoBST(b *testing.B) {
	c, err := NewCrawler(TestDB, &containers.BinaryTree{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = 2000
	c.MaxConcurrentRequests = 1000
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}
