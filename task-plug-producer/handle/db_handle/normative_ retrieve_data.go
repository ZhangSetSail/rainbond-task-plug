package db_handle

import (
	"github.com/goodrain/rainbond-task-plug/db/model"
	"github.com/sirupsen/logrus"
)

func (d *DBAction) RetrieveNormativeData(serviceID string, serviceList []string) ([]*model.ComponentReport, error) {
	logrus.Infof("retrieve normative data")
	var normative []*model.ComponentReport
	if serviceID != "" {
		logrus.Infof("通过 service id 查询")
		d.db.Debug().Find(&normative, "component_id=?", serviceID)
	}
	if len(serviceList) > 0 {
		logrus.Infof("通过 service id list 查询, %v", serviceList)
		d.db.Debug().Where("component_id in (?)", serviceList).Find(&normative)
	}
	return normative, nil
}
