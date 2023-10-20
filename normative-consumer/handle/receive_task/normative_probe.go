package receive_task

import (
	db_model "github.com/goodrain/rainbond-task-plug/db/model"
	"github.com/goodrain/rainbond-task-plug/db/mysql"
	"github.com/goodrain/rainbond-task-plug/db/rainbond_model"
	"github.com/goodrain/rainbond-task-plug/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type ProbeNormative struct {
	DB *gorm.DB
}

func (s ProbeNormative) Check(ni model.NormativeInspectionModel) {
	var probe rainbond_model.TenantServiceProbe
	err := s.DB.First(&probe, "service_id=?", ni.ComponentID).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		probeReport := db_model.ComponentReport{
			CreateTime:  time.Now(),
			Level:       1,
			Message:     "组件未设置健康检测",
			ComponentID: ni.ComponentID,
			PrimaryLink: "",
			Type:        "normative",
		}
		err := s.DB.Debug().Create(probeReport).Error
		if err != nil {
			logrus.Errorf("create service normative probe record failure: %v", err)
		}
	}
}

func NewProbeNormative() *ProbeNormative {
	db := mysql.GetDB()
	return &ProbeNormative{
		DB: db,
	}
}
