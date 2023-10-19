package main

import (
	"fmt"
	"github.com/goodrain/rainbond-task-plug/cmd/normative-consumer/config"
	"github.com/goodrain/rainbond-task-plug/cmd/normative-consumer/server"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	config.Parse()
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)
	switch config.GetNormativeConsumerServer().LogLevel {
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func main() {
	if err := server.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
