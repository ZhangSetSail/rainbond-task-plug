package dispatch_tasks

import (
	"encoding/json"
	"github.com/goodrain/rainbond-safety/model"
	"github.com/sirupsen/logrus"
)

func (m *ManagerDispatchTasks) CreateSourceCodeInspectionTask() error {
	cdm := model.CodeDetectionModel{
		ProjectName:   "zqh-test",
		RepositoryURL: "https://github.com/SonarSource/sonar-scanner-cli-docker.git",
	}
	cdmJsonByte, err := json.Marshal(cdm)
	if err != nil {
		logrus.Errorf("有报错%v", err)
		return err
	}
	err = m.nc.Publish(m.config.Subscribe, cdmJsonByte)
	if err != nil {
		logrus.Errorf("有报错%v", err)
		return err
	}
	return nil
}
