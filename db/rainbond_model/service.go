package rainbond_model

type TenantServiceVolume struct {
	ServiceID  string `gorm:"column:service_id;type:varchar(32);comment:'组件ID'" json:"service_id"`
	VolumeType string `gorm:"column:volume_type;size:64" json:"volume_type"`
}

// TableName 表名
func (t *TenantServiceVolume) TableName() string {
	return "tenant_service_volume"
}

type TenantServiceENV struct {
	ServiceID string `gorm:"column:service_id;type:varchar(32);comment:'组件ID'" json:"service_id"`
	AttrName  string `gorm:"column:attr_name;size:1024" validate:"env_name|required" json:"attr_name"`
	AttrValue string `gorm:"column:attr_value;type:text" validate:"env_value|required" json:"attr_value"`
}

// TableName 表名
func (t *TenantServiceENV) TableName() string {
	return "tenant_service_env_var"
}

type TenantServiceAttributes struct {
	ComponentID string `gorm:"column:component_id;type:varchar(32)" json:"component_id"`
	Name        string `gorm:"column:name" json:"name"`
}

func (t *TenantServiceAttributes) TableName() string {
	return "component_k8s_attributes"
}

type TenantServiceProbe struct {
	ServiceID string `gorm:"column:service_id;type:varchar(32);comment:'组件ID'" json:"service_id"`
}

func (t *TenantServiceProbe) TableName() string {
	return "service_probe"
}
