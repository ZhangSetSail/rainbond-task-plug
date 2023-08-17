package handle

import (
	"context"
	"github.com/goodrain/rainbond-safety/cmd/safety-consumer/option"
	"github.com/goodrain/rainbond-safety/safety-consumer/handle/receive_task"
	"github.com/nats-io/nats.go"
)

func InitHandle(ctx context.Context, nc *nats.Conn, config option.Config) {
	defaultManagerReceiveTasks = receive_task.CreateManagerReceiveTask(ctx, nc, config)
}

var defaultManagerReceiveTasks *receive_task.ManagerReceiveTask

func GetManagerReceiveTasks() *receive_task.ManagerReceiveTask {
	return defaultManagerReceiveTasks
}
