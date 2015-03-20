// The Mir Sprinter command.
package main

import (
	"flag"
	"fmt"
	"log"
	"mir/sprinter"
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
	c, err := sprinter.NewCrawler()
	if err != nil {
		panic(err)
	}
	err = c.AddURL("http://www.google.com")
	if err != nil {
		panic(err)
	}
	result, err := c.RetrieveHTML(0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", result)
}
