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

type ENVNormative struct {
	DB *gorm.DB
}

func (s ENVNormative) Check(ni model.NormativeInspectionModel) {
	var envs []rainbond_model.TenantServiceENV
	s.DB.Find(&envs, "service_id=?", ni.ComponentID)
	var componentReportList []*db_model.ComponentReport
	if envs != nil && len(envs) >= 0 {
		for _, env := range envs {
			var message string
			switch env.AttrName {
			case "ES_SELECTNODE":
				message = "组件使用了环境变量 ES_SELECTNODE 节点选择属性，发布安装后会影响组件正常使用。"
			case "PAUSE":
				message = "组件使用了环境变量 PAUSE 容器暂停，非正常业务运行。"
			}
			if message != "" {
				componentReportList = append(componentReportList, &db_model.ComponentReport{
					CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
					Level:       1,
					Message:     message,
					ComponentID: ni.ComponentID,
					PrimaryLink: "",
					Type:        "normative",
				})
			}
		}
		err := s.DB.Debug().Create(&componentReportList).Error
		if err != nil {
			logrus.Errorf("create service normative env record failure: %v", err)
		}
	}
}

func NewENVNormative() *ENVNormative {
	db := mysql.GetDB()
	return &ENVNormative{
		DB: db,
	}
}
