package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/danlin/mutago/parser"
)

func main() {
	flag.Parse()
	path := flag.Arg(0)
	if path == "" {
		fmt.Printf("Usage: mutago path")
		return
	}
	start := time.Now()
	parser.Parse(path)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s", elapsed)
}
