package main

import (
	"flag"
	"fmt"

	"github.com/danlin/mutago/parser"
)

func main() {
	flag.Parse()
	path := flag.Arg(0)
	if path == "" {
		fmt.Printf("Usage: mutago path")
		return
	}
	parser.Parse(path)
}
