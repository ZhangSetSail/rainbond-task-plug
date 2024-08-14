package handle

import (
	"context"
	"github.com/goodrain/rainbond-task-plug/cmd/safety-consumer/config"
	"github.com/goodrain/rainbond-task-plug/db/mysql"
	"github.com/goodrain/rainbond-task-plug/es/es"
	"github.com/goodrain/rainbond-task-plug/safety-consumer/handle/receive_task"
	"github.com/nats-io/nats.go"
)

func InitHandle(ctx context.Context, nc *nats.Conn, config config.Config) {
	db := mysql.GetDB()
	esc := es.GetES()
	defaultManagerReceiveTasks = receive_task.CreateManagerReceiveTask(ctx, nc, config, db, esc)
}

var defaultManagerReceiveTasks *receive_task.ManagerReceiveTask

func GetManagerReceiveTasks() *receive_task.ManagerReceiveTask {
	return defaultManagerReceiveTasks
}
