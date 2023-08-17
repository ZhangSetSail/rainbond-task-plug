package option

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

// Config config
type Config struct {
	NatsAPI   string
	Subscribe string
}

// ProducerServer  apiserver server
type ProducerServer struct {
	Config
	LogLevel string
}

// NewProducerServer new server
func NewProducerServer() *ProducerServer {
	return &ProducerServer{}
}

// AddFlags config
func (a *ProducerServer) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&a.LogLevel, "log-level", "info", "log level")
	fs.StringVar(&a.NatsAPI, "nats-api", "47.93.219.143:10007", "nats host:127.0.0.1:4222")
	fs.StringVar(&a.Subscribe, "subscribe", "rainbond", "subscription number name")
}

// SetLog 设置log
func (a *ProducerServer) SetLog() {
	level, err := logrus.ParseLevel(a.LogLevel)
	if err != nil {
		fmt.Println("set log level error." + err.Error())
		return
	}
	logrus.SetLevel(level)
}
