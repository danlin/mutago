package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/danlin/mutago/backend"
	"github.com/danlin/mutago/parser"
)

func main() {
	b, err := backend.Open("./data")
	if err != nil {
		panic(err)
	}
	defer b.Close()

	flag.Parse()
	path := flag.Arg(0)
	if path == "" {
		fmt.Print("Usage: mutago path")
		return
	}
	start := time.Now()
	parser.Parse(path)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s", elapsed)
}
