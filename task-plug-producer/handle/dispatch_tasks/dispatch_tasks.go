package dispatch_tasks

import (
	"context"
	"github.com/nats-io/nats.go"
)

type DispatchTasksAction struct {
	ctx context.Context
	nc  *nats.Conn
}

func CreateDispatchTasksHandle(ctx context.Context, nc *nats.Conn) *DispatchTasksAction {
	return &DispatchTasksAction{
		ctx: ctx,
		nc:  nc,
	}
}
