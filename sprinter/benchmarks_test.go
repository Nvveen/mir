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
	c.MaxRequests = 1000
	c.MaxConcurrentRequests = 1
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}
