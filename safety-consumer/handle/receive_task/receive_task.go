package receive_task

import (
	"context"
	"github.com/goodrain/rainbond-task-plug/cmd/safety-consumer/config"
	"github.com/goodrain/rainbond-task-plug/es/es"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type ManagerReceiveTask struct {
	ctx    context.Context
	nc     *nats.Conn
	config config.Config
	db     *gorm.DB
	esCli  *es.ComponentReportRepo
}

func CreateManagerReceiveTask(ctx context.Context, nc *nats.Conn, config config.Config, db *gorm.DB, esc *es.ComponentReportRepo) *ManagerReceiveTask {
	return &ManagerReceiveTask{
		ctx:    ctx,
		nc:     nc,
		config: config,
		db:     db,
		esCli:  esc,
	}
}
