package handle

import (
	"github.com/goodrain/rainbond-task-plug/normative-consumer/handle/receive_task"
	"github.com/goodrain/rainbond-task-plug/pkg"
	"github.com/sirupsen/logrus"
)

func InitHandle() {
	logrus.Infof("init handle")
	nc := pkg.GetNatsClient()
	ctx := pkg.GetCTX()
	defaultManagerReceiveTasks = receive_task.CreateManagerReceiveTask(ctx, nc)
}

var defaultManagerReceiveTasks *receive_task.ManagerReceiveTask

func GetManagerReceiveTasks() *receive_task.ManagerReceiveTask {
	return defaultManagerReceiveTasks
}
