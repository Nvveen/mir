// The sprinter package is a concurrent webquery engine.
package sprinter

import (
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

func init() {
	// Setup loggers
	Info = log.New(os.Stdout, "INFO ", log.LstdFlags)
	Error = log.New(os.Stderr, "ERROR ", log.LstdFlags)
}
