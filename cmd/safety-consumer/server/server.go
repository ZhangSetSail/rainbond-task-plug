package server

import (
	"context"
	"github.com/goodrain/rainbond-task-plug/cmd/safety-consumer/config"
	"github.com/goodrain/rainbond-task-plug/db/mysql"
	"github.com/goodrain/rainbond-task-plug/safety-consumer/api_router"
	"github.com/goodrain/rainbond-task-plug/safety-consumer/handle"
	nats "github.com/nats-io/nats.go"
)

func Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c := config.GetSafetyConsumerServer()
	natsAddr := c.NatsHost + ":" + c.NatsPort
	nc, err := nats.Connect(natsAddr)
	if err != nil {
		return err
	}
	err = mysql.InitDB(c.DB)
	if err != nil {
		return err
	}
	handle.InitHandle(ctx, nc, c.Config)
	err = handle.GetManagerReceiveTasks().DigestionSourceCodeInspectionTask()
	if err != nil {
		return err
	}
	return api_router.InitRouter()
}
