package receive_task

import (
	"encoding/json"
	"github.com/goodrain/rainbond-safety/model"
	"github.com/goodrain/rainbond-safety/safety-consumer/handle/clone"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"path"
)

func (t *ManagerReceiveTask) DigestionSourceCodeInspectionTask() error {
	logrus.Infof("begion receive source code inspection task")
	_, err := t.nc.QueueSubscribe(t.config.Subscribe, t.config.SubscribeQueue, func(m *nats.Msg) {
		var cdm model.CodeDetectionModel
		err := json.Unmarshal(m.Data, &cdm)
		if err != nil {
			logrus.Errorf("json unmarshal failure: %v", err)
			return
		}
		codePath := path.Join(t.config.CodeStoragePath, cdm.ProjectName)
		_, _, err = clone.GitClone(cdm, codePath, 10, t.ctx)
		if err != nil {
			logrus.Errorf("git clone failure: %v", err)
			return
		}

	})
	return err
}
