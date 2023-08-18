package dispatch_tasks

import (
	"encoding/json"
	"github.com/goodrain/rainbond-safety/model"
	"github.com/sirupsen/logrus"
)

func (m *ManagerDispatchTasks) CreateSourceCodeInspectionTask() error {
	cdm := model.CodeDetectionModel{
		ProjectName:   "new-test",
		RepositoryURL: "https://github.com/goodrain/rainbond.git",
	}
	cdmJsonByte, err := json.Marshal(cdm)
	if err != nil {
		logrus.Errorf("json marshal code detection model failure: %v", err)
		return err
	}
	err = m.nc.Publish(m.config.Subscribe, cdmJsonByte)
	if err != nil {
		logrus.Errorf("publish code detection model failure: %v", err)
		return err
	}
	return nil
}
