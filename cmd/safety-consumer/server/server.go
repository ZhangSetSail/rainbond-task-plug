package server

import (
	"context"
	"github.com/goodrain/rainbond-safety/cmd/safety-consumer/option"
	"github.com/goodrain/rainbond-safety/safety-consumer/api_router"
	"github.com/goodrain/rainbond-safety/safety-consumer/handle"
	nats "github.com/nats-io/nats.go"
)

func Run(s *option.ConsumerServer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nc, err := nats.Connect(s.NatsAPI)
	if err != nil {
		return err
	}
	handle.InitHandle(ctx, nc, s.Config)
	err = handle.GetManagerReceiveTasks().DigestionSourceCodeInspectionTask()
	if err != nil {
		return err
	}
	return api_router.InitRouter()
}
