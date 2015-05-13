// The Mir Sprinter command.
package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"syscall"
	"testing"

	"github.com/Nvveen/mir/containers"
	"github.com/Nvveen/mir/sink"
	"github.com/Nvveen/mir/sprinter"
	"github.com/Nvveen/mir/storage"
)

var (
	uri     = flag.String("uri", "", "Pass an initial URI to start crawling from")
	bench   = flag.Bool("bench", false, "Instead of actually crawling, benchmark from the testing suite")
	verbose = flag.Bool("verbose", false, "Show output of the crawling")
)

func runBench() {
	// Set maximum number of CPUs
	// TODO panics?
	// runtime.GOMAXPROCS(runtime.NumCPU())
	// Start local server
	go sink.RunSink()
	BenchVerbose = *verbose
	// Benchmark
	f := func(bn func(*testing.B)) {
		name := runtime.FuncForPC(reflect.ValueOf(bn).Pointer()).Name()
		for i := 0; i < 5; i++ {
			log.Printf("Benchmarking (%d) %s...\n", i, name)
			b := testing.Benchmark(bn)
			log.Printf("\tAllocated bytes per operation: %d\n", b.AllocsPerOp())
			log.Printf("\tAllocations per operation: %d\n", b.AllocedBytesPerOp())
			log.Printf("\tNanoseconds per operation: %d\n", b.NsPerOp())
			log.Printf("\tMilliseconds per operation: %d\n", b.NsPerOp()/1000000)
			log.Printf("\tSeconds per operation: %d\n", b.NsPerOp()/1000000000)
		}
	}
	f(BenchmarkCrawl_SequentialStaticList)
	f(BenchmarkCrawl_SequentialStaticBST)
	f(BenchmarkCrawl_Concurrent10StaticList)
	f(BenchmarkCrawl_Concurrent50StaticList)
	f(BenchmarkCrawl_Concurrent100StaticList)
	f(BenchmarkCrawl_Concurrent500StaticList)
	f(BenchmarkCrawl_Concurrent1000StaticList)
	f(BenchmarkCrawl_Concurrent100StaticBST)
	// start mongo
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
			err := storage.TestDB.StopMongoTesting()
			if err != nil {
				log.Println(err)
			}
			log.Println(r)
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
	f(BenchmarkCrawl_Concurrent100MongoList)
	f(BenchmarkCrawl_Concurrent100MongoBST)
	err = storage.TestDB.StopMongoTesting()
	if err != nil {
		panic(err)
	}
	os.Exit(0)
}

func init() {
	// Parse flags
	flag.Parse()
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln("Unrecoverable:", err)
		}
	}()
	// Parse flags
	flag.Parse()
	if *bench {
		runBench()
	}
	if uri == nil || len(*uri) == 0 {
		panic("crawler needs a URI to crawl if not benchmarking.")
	}
	c, err := sprinter.NewCrawler(storage.NewMockStorage(), &containers.List{})
	if err != nil {
		panic(err)
	}
	c.Verbose = *verbose
	err = c.Crawl(*uri)
	if err != nil {
		panic(err)
	}
}
