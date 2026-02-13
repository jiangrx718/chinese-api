package model

import "time"

// CRMRole 角色表
type CRMRole struct {
	Id        int       `gorm:"column:id;type:bigint;not null;primaryKey;autoIncrement;comment:主键;" json:"id"`
	RoleId    string    `gorm:"column:role_id;type:char(36);not null;default:'';index:idx_role_id;comment:业务主键;" json:"role_id"`
	RoleName  string    `gorm:"column:role_name;type:varchar(128);not null;default:'';comment:角色名称" json:"role_name"`
	Status    string    `gorm:"column:status;type:varchar(32);not null;default:'on';comment:状态,off禁用,on启用" json:"status"`
	IsInit    string    `gorm:"column:is_init;type:varchar(32);not null;default:'off';comment:是否初始数据,off否,on是" json:"is_init"`
	IsSuper   string    `gorm:"column:is_super;type:varchar(32);not null;default:'off';comment:是否超管数据,off否,on是" json:"is_super"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;index:idx_created_at;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
}

func (r *CRMRole) TableName() string {
	return "crm_role"
}

// CRMRolePermission 角色权限表
type CRMRolePermission struct {
	Id           int       `gorm:"column:id;type:bigint;not null;primaryKey;autoIncrement;comment:自增主键;" json:"id"`
	RoleId       string    `gorm:"column:role_id;type:char(36);not null;default:'';index:idx_role_id;comment:角色ID;" json:"role_id"`
	PermissionId string    `gorm:"column:permission_id;type:char(36);not null;default:'';comment:权限ID" json:"permission_id"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;index:idx_created_at;comment:创建时间" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
}

func (r *CRMRolePermission) TableName() string {
	return "crm_role_permission"
}
