package handle

import (
	"context"
	"github.com/goodrain/rainbond-safety/safety-producer/handle/dispatch_tasks"
	"github.com/nats-io/nats.go"
	"k8s.io/client-go/rest"
)

func InitHandle(ctx context.Context, config *rest.Config, nc *nats.Conn) {

	defaultManagerClientGo = dispatch_tasks.CreateManagerDispatchTasks(ctx, nc)
}

var defaultManagerClientGo *dispatch_tasks.ManagerDispatchTasks

func GetManagerClientGo() *dispatch_tasks.ManagerDispatchTasks {
	return defaultManagerClientGo
}
