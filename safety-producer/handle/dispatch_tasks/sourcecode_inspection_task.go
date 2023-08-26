package dispatch_tasks

import (
	"encoding/json"
	"github.com/goodrain/rainbond-safety/model"
	"github.com/sirupsen/logrus"
)

func (m *ManagerDispatchTasks) CreateSourceCodeInspectionTask(projectName, url string) error {
	logrus.Infof(projectName)
	cdm := model.CodeDetectionModel{
		ProjectName:   projectName,
		RepositoryURL: url,
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
