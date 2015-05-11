// The sink is a program that acts as a testing backend to the crawler,
// so we don't have to pester the internet.
package main

import (
	"github.com/Nvveen/mir/sink"
)

func main() {
	sink.RunSink()
}
