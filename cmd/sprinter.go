// The Mir Sprinter command.
package main

import (
	"flag"
	"log"
)

func init() {
	// Parse flags
	flag.Parse()
}

func main() {
	if len(flag.Args()) != 1 {
		log.Fatalf("Invalid number of arguments: %d\n", len(flag.Args()))
	}
	// Parse flags
	flag.Parse()
}
