package model

import "time"

type ComponentReport struct {
	PrimaryLink string    `gorm:"column:primary_link;comment:'路由'" json:"primary_link"`
	Level       int       `gorm:"column:level;comment:'报警等级'" json:"level"`
	Message     string    `gorm:"column:message;comment:'报警信息'" json:"message"`
	ComponentID string    `gorm:"column:component_id;type:varchar(32);comment:'组件ID'" json:"component_id"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time;comment:'创建时间'"`
	Type        string    `json:"type" gorm:"cloumn:type;comment:'报告类型'"`
}

func (t *ComponentReport) TableName() string {
	return "component_report"
}
