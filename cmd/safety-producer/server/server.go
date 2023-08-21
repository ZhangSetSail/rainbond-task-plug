package server

import (
	"context"
	"github.com/goodrain/rainbond-safety/cmd/safety-producer/option"
	"github.com/goodrain/rainbond-safety/safety-producer/api_router"
	"github.com/goodrain/rainbond-safety/safety-producer/handle"
	init_watch "github.com/goodrain/rainbond-safety/safety-producer/handle/k8s-watch/init-watch"
	"github.com/goodrain/rainbond-safety/util"
	nats "github.com/nats-io/nats.go"
)

func Run(s *option.ProducerServer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	clientSet, _, err := util.InitK8SClient()
	if err != nil {
		return err
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
