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
	c.MaxRequests = MaxRequests
	c.MaxConcurrentRequests = 1
	c.Verbose = BenchVerbose
	b.ResetTimer()
	c.CrawlSequential("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent10StaticList(b *testing.B) {
	c, err := sprinter.NewCrawler(storage.NewMockStorage(), &containers.List{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = MaxRequests
	c.MaxConcurrentRequests = 10
	c.Verbose = BenchVerbose
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent100StaticList(b *testing.B) {
	c, err := sprinter.NewCrawler(storage.NewMockStorage(), &containers.List{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = MaxRequests
	c.MaxConcurrentRequests = 100
	c.Verbose = BenchVerbose
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent1000StaticList(b *testing.B) {
	c, err := sprinter.NewCrawler(storage.NewMockStorage(), &containers.List{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = MaxRequests
	c.MaxConcurrentRequests = 1000
	c.Verbose = BenchVerbose
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent1000MongoList(b *testing.B) {
	c, err := sprinter.NewCrawler(storage.TestDB, &containers.List{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = MaxRequests
	c.MaxConcurrentRequests = MaxConcurrentRequests
	c.Verbose = BenchVerbose
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent1000MongoBST(b *testing.B) {
	c, err := sprinter.NewCrawler(storage.TestDB, &containers.BinaryTree{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = MaxRequests
	c.MaxConcurrentRequests = MaxConcurrentRequests
	c.Verbose = BenchVerbose
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}
