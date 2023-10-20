package receive_task

import (
	"fmt"
	db_model "github.com/goodrain/rainbond-task-plug/db/model"
	"github.com/goodrain/rainbond-task-plug/db/mysql"
	"github.com/goodrain/rainbond-task-plug/db/rainbond_model"
	"github.com/goodrain/rainbond-task-plug/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type AttributesNormative struct {
	DB *gorm.DB
}

func (s AttributesNormative) Check(ni model.NormativeInspectionModel) {
	var attributes []rainbond_model.TenantServiceAttributes
	s.DB.Find(&attributes, "component_id=?", ni.ComponentID)
	var componentReportList []*db_model.ComponentReport
	if attributes != nil && len(attributes) >= 0 {
		for _, attribute := range attributes {
			var message string
			switch attribute.Name {
			case "volumeClaimTemplate":
				message = "volumeClaimTemplate"
			case "volumes":
				message = "volumes"
			case "volumeMounts":
				message = "volumeMounts"
			case "affinity":
				message = "affinity"
			case "nodeSelector":
				message = "nodeSelector"
			}
			if message != "" {
				componentReportList = append(componentReportList, &db_model.ComponentReport{
					CreateTime:  time.Now(),
					Level:       1,
					Message:     fmt.Sprintf("组件使用了 %s 高级属性，可能会影响组件的发布后的正常安装", message),
					ComponentID: ni.ComponentID,
					PrimaryLink: "",
					Type:        "normative",
				})
			}

		}
		err := s.DB.Debug().Create(&attributes).Error
		if err != nil {
			logrus.Errorf("create service normative attributes record failure: %v", err)
		}
	}
}

func NewAttributesNormative() *AttributesNormative {
	db := mysql.GetDB()
	return &AttributesNormative{
		DB: db,
	}
}
