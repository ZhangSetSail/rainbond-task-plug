package server

import (
	"context"
	"github.com/goodrain/rainbond-task-plug/cmd/safety-consumer/option"
	"github.com/goodrain/rainbond-task-plug/safety-consumer/api_router"
	"github.com/goodrain/rainbond-task-plug/safety-consumer/handle"
	nats "github.com/nats-io/nats.go"
	"os"
)

func Run(s *option.ConsumerServer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if s.NatsAPI == "" {
		s.NatsAPI = os.Getenv("NATS_HOST") + ":" + os.Getenv("NATS_PORT")
	}
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
