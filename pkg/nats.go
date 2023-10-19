package pkg

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

var (
	nc *nats.Conn
)

func InitNatsCli(natsAddr string) error {
	logrus.Infof("init nats cli")
	var err error
	nc, err = nats.Connect(natsAddr)
	return err
}

func GetNatsClient() *nats.Conn {
	return nc
}
