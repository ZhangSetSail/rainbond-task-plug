package rainbond_model

type TenantServiceVolume struct {
	ServiceID  string `gorm:"column:service_id;type:varchar(32);comment:'组件ID'" json:"service_id"`
	VolumeType string `gorm:"column:volume_type;size:64" json:"volume_type"`
}

// TableName 表名
func (t *TenantServiceVolume) TableName() string {
	return "tenant_service_volume"
}
