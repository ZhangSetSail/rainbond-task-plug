package main

import (
	"fmt"
	"github.com/goodrain/rainbond-safety/cmd/safety-producer/server"
	"os"
)

func main() {
	if err := server.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
