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
	SonarToken      string
	SonarHostUrl    string
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
	fs.StringVar(&a.CodeStoragePath, "code-storage-path", "/usr/src/", "code storage address")
	fs.StringVar(&a.SonarToken, "sonar-token", "squ_13a8847dc673ff817ceb82a4a1f41ce4d1d582ab", "sonar token")
	fs.StringVar(&a.SonarHostUrl, "sonar-host-url", "http://39.105.56.19:32097", "sonar host url")
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
