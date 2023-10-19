package receive_task

import (
	"github.com/goodrain/rainbond-task-plug/db/mysql"
	"github.com/goodrain/rainbond-task-plug/model"
	"gorm.io/gorm"
)

type ENVNormative struct {
	DB *gorm.DB
}

func (s ENVNormative) Check(ni model.NormativeInspectionModel) {

}

func NewENVNormative() *ENVNormative {
	db := mysql.GetDB()
	return &ENVNormative{
		DB: db,
	}
}
