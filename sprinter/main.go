// The sprinter package is a concurrent webquery engine.
package main

import (
	"flag"
	"log"
	"os"
)

// TODO split into cmd/ and lib package

var (
	Info  *log.Logger
	Error *log.Logger
)

func init() {
	// Setup loggers
	Info = log.New(os.Stdout, "INFO ", log.LstdFlags)
	Error = log.New(os.Stderr, "ERROR ", log.LstdFlags)
	// Parse flags
	flag.Parse()
}

func main() {
	if len(flag.Args()) != 1 {
		log.Fatalf("Invalid number of arguments: %d\n", len(flag.Args()))
	}
}
