package model

import (
	"time"
)

// CRMPermission 权限菜单表
type CRMPermission struct {
	Id             int       `gorm:"column:id;type:bigint;not null;primaryKey;autoIncrement;comment:主键;" json:"id"`
	PermissionId   string    `gorm:"column:permission_id;type:char(36);not null;default:'';index:idx_permission_id;;comment:业务主键;" json:"permission_id"`
	PermissionName string    `gorm:"column:permission_name;type:varchar(128);not null;default:'';comment:权限名称" json:"permission_name"`
	PermissionURL  string    `gorm:"column:permission_url;type:varchar(128);not null;default:'';comment:权限URL" json:"permission_url"`
	ParentId       string    `gorm:"column:parent_id;type:char(36);not null;default:'';comment:父级id;" json:"parent_id"`
	Status         string    `gorm:"column:status;type:varchar(32);not null;default:'on';comment:状态,off禁用,on启用" json:"status"`
	IsInit         string    `gorm:"column:is_init;type:varchar(32);not null;default:'off';comment:是否初始数据,off否,on是" json:"is_init"`
	PermissionType int       `gorm:"column:permission_type;type:int;not null;default:0;comment:权限类型,1菜单,2按钮,3接口" json:"permission_type"`
	Position       int       `gorm:"column:position;type:int;not null;default:0;comment:排序" json:"position"`
	CreatedAt      time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;index:idx_created_at;comment:创建时间" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
}

func (r *CRMPermission) TableName() string {
	return "crm_permission"
}
