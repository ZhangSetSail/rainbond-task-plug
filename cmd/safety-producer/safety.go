package main

import (
	"fmt"
	"github.com/goodrain/rainbond-safety/cmd/safety-producer/option"
	"github.com/goodrain/rainbond-safety/cmd/safety-producer/server"
	"github.com/spf13/pflag"
	"os"
)

func main() {
	s := option.NewProducerServer()
	s.AddFlags(pflag.CommandLine)
	pflag.Parse()
	s.SetLog()
	if err := server.Run(s); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
