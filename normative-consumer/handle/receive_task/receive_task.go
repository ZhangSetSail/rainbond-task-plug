package receive_task

import (
	"context"
	"github.com/nats-io/nats.go"
)

type ManagerReceiveTask struct {
	ctx context.Context
	nc  *nats.Conn
}

func CreateManagerReceiveTask(ctx context.Context, nc *nats.Conn) *ManagerReceiveTask {
	return &ManagerReceiveTask{
		ctx: ctx,
		nc:  nc,
	}
}
