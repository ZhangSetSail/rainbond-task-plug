package receive_task

import (
	db_model "github.com/goodrain/rainbond-task-plug/db/model"
	"github.com/goodrain/rainbond-task-plug/db/mysql"
	"github.com/goodrain/rainbond-task-plug/db/rainbond_model"
	"github.com/goodrain/rainbond-task-plug/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
	"time"
)

type StorageNormative struct {
	DB *gorm.DB
}

func (s StorageNormative) Check(ni model.NormativeInspectionModel) {
	var volumes []rainbond_model.TenantServiceVolume
	s.DB.Where("volume_type <> ?", "config-file").Find(&volumes, "service_id=?", ni.ComponentID)
	if len(volumes) > 0 {
		logrus.Infof("volumes: %v", volumes)
		records := db_model.ComponentReport{
			CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
			Level:       1,
			Message:     "组件挂载了存储，发布后安装存储数据不会携带，可能会影响组件正常使用",
			ComponentID: ni.ComponentID,
			PrimaryLink: "",
			Type:        "normative",
		}
		err := s.DB.Debug().Create(&records).Error
		if err != nil {
			logrus.Errorf("create service normative voluem record failure: %v", err)
		}
	}
	if len(volumes) == 0 && strings.HasPrefix(ni.ExtendMethod, "state_") {
		records := db_model.ComponentReport{
			CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
			Level:       1,
			Message:     "有状态组件没有挂载存储",
			ComponentID: ni.ComponentID,
			PrimaryLink: "",
			Type:        "normative",
		}
		err := s.DB.Debug().Create(&records).Error
		if err != nil {
			logrus.Errorf("create service normative record failure: %v", err)
		}
	}
}

func NewStorageNormative() *StorageNormative {
	db := mysql.GetDB()
	return &StorageNormative{
		DB: db,
	}
}
