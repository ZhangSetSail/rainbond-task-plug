package handle

import (
	"github.com/goodrain/rainbond-task-plug/db/mysql"
	"github.com/goodrain/rainbond-task-plug/pkg"
	"github.com/goodrain/rainbond-task-plug/task-plug-producer/handle/db_handle"
	"github.com/goodrain/rainbond-task-plug/task-plug-producer/handle/dispatch_tasks"
	"github.com/sirupsen/logrus"
)

func InitHandle() {
	logrus.Infof("init handle")
	nc := pkg.GetNatsClient()
	ctx := pkg.GetCTX()
	db := mysql.GetDB()
	dispatchTasksHandle = dispatch_tasks.CreateDispatchTasksHandle(ctx, nc)
	dbHandle = db_handle.CreateDBHandle(ctx, db)
}

var dispatchTasksHandle *dispatch_tasks.DispatchTasksAction
var dbHandle *db_handle.DBAction

func GetDispatchTasksHandle() *dispatch_tasks.DispatchTasksAction {
	return dispatchTasksHandle
}

func GetDBHandle() *db_handle.DBAction {
	return dbHandle
}
