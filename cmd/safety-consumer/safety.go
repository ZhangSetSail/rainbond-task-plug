package main

import (
	"fmt"
	"github.com/goodrain/rainbond-task-plug/cmd/safety-consumer/option"
	"github.com/goodrain/rainbond-task-plug/cmd/safety-consumer/server"
	"github.com/spf13/pflag"
	"os"
)

func main() {
	s := option.NewConsumerServer()
	s.AddFlags(pflag.CommandLine)
	pflag.Parse()
	s.SetLog()
	if err := server.Run(s); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
