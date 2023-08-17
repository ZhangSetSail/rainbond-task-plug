package handle

import (
	"context"
	"github.com/goodrain/rainbond-safety/cmd/safety-producer/option"
	"github.com/goodrain/rainbond-safety/safety-producer/handle/dispatch_tasks"
	"github.com/nats-io/nats.go"
)

func InitHandle(ctx context.Context, nc *nats.Conn, config option.Config) {
	defaultManagerDispatchTasks = dispatch_tasks.CreateManagerDispatchTasks(ctx, nc, config)
}

var defaultManagerDispatchTasks *dispatch_tasks.ManagerDispatchTasks

func GetManagerDispatchTasks() *dispatch_tasks.ManagerDispatchTasks {
	return defaultManagerDispatchTasks
}
