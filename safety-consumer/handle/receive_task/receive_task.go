package receive_task

import (
	"context"
	"github.com/goodrain/rainbond-task-plug/cmd/safety-consumer/option"
	"github.com/nats-io/nats.go"
)

type ManagerReceiveTask struct {
	ctx    context.Context
	nc     *nats.Conn
	config option.Config
}

func CreateManagerReceiveTask(ctx context.Context, nc *nats.Conn, config option.Config) *ManagerReceiveTask {
	return &ManagerReceiveTask{
		ctx:    ctx,
		nc:     nc,
		config: config,
	}
}
