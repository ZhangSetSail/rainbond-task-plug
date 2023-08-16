package dispatch_tasks

import (
	"encoding/json"
	"github.com/goodrain/rainbond-safety/model"
	"github.com/sirupsen/logrus"
)

func (m *ManagerDispatchTasks) CreateTask() {
	cdm := model.CodeDetectionModel{
		ProjectName:   "zqh-test",
		SourceAddress: "abc",
	}
	cdmJsonByte, err := json.Marshal(cdm)
	if err != nil {
		logrus.Errorf("有报错%v", err)
		return
	}
	err = m.nc.Publish("foo", cdmJsonByte)
	if err != nil {
		logrus.Errorf("有报错%v", err)
	}
}
