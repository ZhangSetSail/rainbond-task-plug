package server

import (
	"context"
	"github.com/goodrain/rainbond-safety/safety-producer/api_router"
	"github.com/goodrain/rainbond-safety/safety-producer/handle"
	init_watch "github.com/goodrain/rainbond-safety/safety-producer/handle/k8s-watch/init-watch"
	nats "github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//设置日志输出格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	clientSet, config, err := InitK8SClient()
	if err != nil {
		return err
	}
	nc, err := nats.Connect("47.93.219.143:10007")
	if err != nil {
		return err
	}

	mw := init_watch.CreateResourceWatch(clientSet)
	mw.Start()

	handle.InitHandle(ctx, config, nc)
	return api_router.InitRouter()
}
