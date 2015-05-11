package sprinter

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/Nvveen/mir/containers"
	"github.com/Nvveen/mir/storage"
)

// Sequential <-> Concurrent
// Static <-> Actual MongoDB backend
// List <-> Binary Search Tree
// Remember to determine the parallelism.

func TestMain(m *testing.M) {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT)
	go func() {
		<-sigc
		storage.TestDB.StopMongoTesting()
		log.Fatal("signal interrupt caught")
	}()
	defer func() {
		if r := recover(); r != nil {
			log.Println("recovering mongo error")
			storage.TestDB.StopMongoTesting()
			log.Print(r)
		}
	}()
	storage.TestDB = storage.NewTestDBMongo(&storage.MongoDB{
		Host: "127.0.0.1",
		Port: "40001",
	})
	err := storage.TestDB.StartMongoTesting()
	if err != nil {
		panic(err)
	}
	// run tests
	ret := m.Run()
	err = storage.TestDB.StopMongoTesting()
	if err != nil {
		panic(err)
	}
	os.Exit(ret)
}

func BenchmarkCrawl_SequentialStaticList(b *testing.B) {
	c, err := NewCrawler(storage.NewMockStorage(), &containers.List{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = 2000
	c.MaxConcurrentRequests = 1
	b.ResetTimer()
	c.CrawlSequential("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent10StaticList(b *testing.B) {
	c, err := NewCrawler(storage.NewMockStorage(), &containers.List{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = 2000
	c.MaxConcurrentRequests = 10
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent100StaticList(b *testing.B) {
	c, err := NewCrawler(storage.NewMockStorage(), &containers.List{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = 2000
	c.MaxConcurrentRequests = 100
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent1000StaticList(b *testing.B) {
	c, err := NewCrawler(storage.NewMockStorage(), &containers.List{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = 2000
	c.MaxConcurrentRequests = 1000
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent1000MongoList(b *testing.B) {
	c, err := NewCrawler(storage.TestDB, &containers.List{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = 2000
	c.MaxConcurrentRequests = 1000
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}

func BenchmarkCrawl_Concurrent1000MongoBST(b *testing.B) {
	c, err := NewCrawler(storage.TestDB, &containers.BinaryTree{})
	if err != nil {
		b.Fatal(err)
	}
	c.MaxRequests = 2000
	c.MaxConcurrentRequests = 1000
	b.ResetTimer()
	c.Crawl("http://localhost:8080")
}
