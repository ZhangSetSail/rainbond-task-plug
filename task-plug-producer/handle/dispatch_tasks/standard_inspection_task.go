package dispatch_tasks

import (
	"encoding/json"
	"github.com/goodrain/rainbond-task-plug/model"
	"github.com/sirupsen/logrus"
)

func (m *DispatchTasksAction) CreateNormativeInspectionTask(componentID, extendMethod string) error {
	logrus.Infof("task: normative inspection")
	cdm := model.NormativeInspectionModel{
		ComponentID:  componentID,
		ExtendMethod: extendMethod,
	}
	cdmJsonByte, err := json.Marshal(cdm)
	if err != nil {
		logrus.Errorf("json marshal normative detection model failure: %v", err)
		return err
	}
	err = m.nc.Publish(model.NORMATIVE_INSPECTION, cdmJsonByte)
	if err != nil {
		logrus.Errorf("publish normative detection model failure: %v", err)
		return err
	}
	return nil
}
