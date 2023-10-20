package receive_task

import (
	db_model "github.com/goodrain/rainbond-task-plug/db/model"
	"github.com/goodrain/rainbond-task-plug/db/mysql"
	"github.com/goodrain/rainbond-task-plug/model"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/json"
	"sync"
)

type NormativeInspectionTask interface {
	Check(ni model.NormativeInspectionModel)
}

var normativeInspectionTaskList []NormativeInspectionTask

func (t *ManagerReceiveTask) DigestionNormativeInspectionTask() error {
	logrus.Infof("begion receive normative inspection task")
	addNormativeInspectionTask()
	_, err := t.nc.QueueSubscribe(model.NORMATIVE_INSPECTION, "", func(m *nats.Msg) {
		var ni model.NormativeInspectionModel
		err := json.Unmarshal(m.Data, &ni)
		if err != nil {
			logrus.Errorf("json unmarshal failure: %v", err)
			return
		}
		db := mysql.GetDB()
		db.Debug().Where("component_id = ? and type = 'normative'", ni.ComponentID).Delete(&db_model.ComponentReport{})
		var wg sync.WaitGroup
		for _, task := range normativeInspectionTaskList {
			wg.Add(1)
			go func(task NormativeInspectionTask) {
				defer wg.Done()
				task.Check(ni)
			}(task)
		}
		wg.Wait()
	})
	return err
}

func addNormativeInspectionTask() {
	normativeInspectionTaskList = append(normativeInspectionTaskList, NewStorageNormative())
	normativeInspectionTaskList = append(normativeInspectionTaskList, NewENVNormative())
	normativeInspectionTaskList = append(normativeInspectionTaskList, NewAttributesNormative())
	normativeInspectionTaskList = append(normativeInspectionTaskList, NewProbeNormative())
	normativeInspectionTaskList = append(normativeInspectionTaskList, NewProcessNormative())
}
