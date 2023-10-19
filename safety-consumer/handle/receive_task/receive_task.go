package receive_task

import (
	"context"
	"github.com/goodrain/rainbond-task-plug/cmd/safety-consumer/config"
	"github.com/jinzhu/gorm"
	"github.com/nats-io/nats.go"
)

type ManagerReceiveTask struct {
	ctx    context.Context
	nc     *nats.Conn
	config config.Config
	db     *gorm.DB
}

func CreateManagerReceiveTask(ctx context.Context, nc *nats.Conn, config config.Config, db *gorm.DB) *ManagerReceiveTask {
	return &ManagerReceiveTask{
		ctx:    ctx,
		nc:     nc,
		config: config,
		db:     db,
	}
}
