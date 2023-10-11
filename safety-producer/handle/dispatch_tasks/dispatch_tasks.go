package dispatch_tasks

import (
	"context"
	"github.com/goodrain/rainbond-task-plug/cmd/task-plug-producer/option"
	"github.com/nats-io/nats.go"
)

type ManagerDispatchTasks struct {
	ctx    context.Context
	nc     *nats.Conn
	config option.Config
}

func CreateManagerDispatchTasks(ctx context.Context, nc *nats.Conn, config option.Config) *ManagerDispatchTasks {
	return &ManagerDispatchTasks{
		ctx:    ctx,
		nc:     nc,
		config: config,
	}
}
