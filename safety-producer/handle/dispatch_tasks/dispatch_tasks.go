package dispatch_tasks

import (
	"context"
	"github.com/nats-io/nats.go"
)

type ManagerDispatchTasks struct {
	ctx context.Context
	nc  *nats.Conn
}

func CreateManagerDispatchTasks(ctx context.Context, nc *nats.Conn) *ManagerDispatchTasks {
	return &ManagerDispatchTasks{
		ctx: ctx,
		nc:  nc,
	}
}
