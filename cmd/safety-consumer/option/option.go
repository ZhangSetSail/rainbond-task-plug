package option

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

// Config config
type Config struct {
	NatsAPI         string
	Subscribe       string
	SubscribeQueue  string
	CodeStoragePath string
}

// ConsumerServer  consumer server
type ConsumerServer struct {
	Config
	LogLevel string
}

// NewConsumerServer new server
func NewConsumerServer() *ConsumerServer {
	return &ConsumerServer{}
}

// AddFlags config
func (a *ConsumerServer) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&a.LogLevel, "log-level", "info", "log level")
	fs.StringVar(&a.NatsAPI, "nats-api", "47.93.219.143:10007", "nats host:127.0.0.1:4222")
	fs.StringVar(&a.Subscribe, "subscribe", "rainbond", "subscription number name")
	fs.StringVar(&a.SubscribeQueue, "subscribe-queue", "rainbond", "subscribe queue name")
	fs.StringVar(&a.CodeStoragePath, "code-storage-path", "/Users/zhangqihang/MyWork/push/github/rainbond-safety", "code storage address")
}

// SetLog 设置log
func (a *ConsumerServer) SetLog() {
	level, err := logrus.ParseLevel(a.LogLevel)
	if err != nil {
		fmt.Println("set log level error." + err.Error())
		return
	}
	logrus.SetLevel(level)
}
