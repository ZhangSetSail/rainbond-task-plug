package server

import (
	"context"
	"github.com/goodrain/rainbond-task-plug/cmd/task-plug-producer/option"
	"github.com/goodrain/rainbond-task-plug/safety-producer/api_router"
	"github.com/goodrain/rainbond-task-plug/safety-producer/handle"
	init_watch "github.com/goodrain/rainbond-task-plug/safety-producer/handle/k8s-watch/init-watch"
	"github.com/goodrain/rainbond-task-plug/util"
	nats "github.com/nats-io/nats.go"
	"os"
)

func Run(s *option.ProducerServer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	clientSet, _, err := util.InitK8SClient()
	if err != nil {
		return err
	}
	if s.NatsAPI == "" {
		s.NatsAPI = os.Getenv("NATS_HOST") + ":" + os.Getenv("NATS_PORT")
	}
	nc, err := nats.Connect(s.NatsAPI)
	if err != nil {
		return err
	}

	mw := init_watch.CreateResourceWatch(clientSet)
	mw.Start()

	handle.InitHandle(ctx, nc, s.Config)
	return api_router.InitRouter()
}