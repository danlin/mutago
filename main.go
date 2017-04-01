package main

import (
	"fmt"
	"os"

	"github.com/danlin/mutago/backend"
	"github.com/danlin/mutago/cmd"
)

var (
	srv *backend.Service
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
