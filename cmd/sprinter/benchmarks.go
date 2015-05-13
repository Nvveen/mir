package main

import (
	"testing"

	"github.com/Nvveen/mir/containers"
	"github.com/Nvveen/mir/sprinter"
	"github.com/Nvveen/mir/storage"
)

// Sequential <-> Concurrent
// Static <-> Actual MongoDB backend
// List <-> Binary Search Tree
// Remember to determine the parallelism.

var (
	BenchVerbose          bool
	MaxRequests           = 2000
	MaxConcurrentRequests = 1000
)

func BenchmarkCrawl_SequentialStaticList(b *testing.B) {
	c, err := sprinter.NewCrawler(storage.NewMockStorage(), &containers.List{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = 2000
	c.MaxConcurrentRequests = 1
	b.ResetTimer()
	c.CrawlSequential("http://localhost:8080")
}

func BenchmarkCrawl_SequentialStaticBST(b *testing.B) {
	c, err := sprinter.NewCrawler(storage.NewMockStorage(), &containers.List{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = 2000
	c.MaxConcurrentRequests = 1
	b.ResetTimer()
	c.CrawlSequential("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent10StaticList(b *testing.B) {
	c, err := sprinter.NewCrawler(storage.NewMockStorage(), &containers.List{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = 2000
	c.MaxConcurrentRequests = 10
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent50StaticList(b *testing.B) {
	c, err := sprinter.NewCrawler(storage.NewMockStorage(), &containers.List{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = 2000
	c.MaxConcurrentRequests = 50
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent100StaticList(b *testing.B) {
	c, err := sprinter.NewCrawler(storage.NewMockStorage(), &containers.List{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = 2000
	c.MaxConcurrentRequests = 100
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent500StaticList(b *testing.B) {
	c, err := sprinter.NewCrawler(storage.NewMockStorage(), &containers.List{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = 2000
	c.MaxConcurrentRequests = 500
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent1000StaticList(b *testing.B) {
	c, err := sprinter.NewCrawler(storage.NewMockStorage(), &containers.List{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = 2000
	c.MaxConcurrentRequests = 1000
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent100StaticBST(b *testing.B) {
	c, err := sprinter.NewCrawler(storage.NewMockStorage(), &containers.BinaryTree{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = 2000
	c.MaxConcurrentRequests = 100
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent100MongoList(b *testing.B) {
	c, err := sprinter.NewCrawler(storage.TestDB, &containers.List{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = 2000
	c.MaxConcurrentRequests = 100
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent100MongoBST(b *testing.B) {
	c, err := sprinter.NewCrawler(storage.TestDB, &containers.BinaryTree{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = 2000
	c.MaxConcurrentRequests = 100
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}
