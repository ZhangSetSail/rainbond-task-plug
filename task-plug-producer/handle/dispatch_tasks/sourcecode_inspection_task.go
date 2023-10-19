package dispatch_tasks

import (
	"encoding/json"
	"github.com/goodrain/rainbond-task-plug/model"
	"github.com/sirupsen/logrus"
)

func (m *DispatchTasksAction) CreateSourceCodeInspectionTask(projectName, url string) error {
	logrus.Infof("task: source code inspection")
	cdm := model.CodeInspectionModel{
		ProjectName:   projectName,
		RepositoryURL: url,
	}
	cdmJsonByte, err := json.Marshal(cdm)
	if err != nil {
		logrus.Errorf("json marshal code detection model failure: %v", err)
		return err
	}
	err = m.nc.Publish(model.SOURCE_CODE_INSPECTION, cdmJsonByte)
	if err != nil {
		logrus.Errorf("publish code detection model failure: %v", err)
		return err
	}
	return nil
}
