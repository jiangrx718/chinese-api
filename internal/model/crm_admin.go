package model

import (
	"time"
)

// CRMAdmin 管理员表
type CRMAdmin struct {
	Id        int       `gorm:"column:id;type:bigint;not null;primaryKey;autoIncrement;comment:主键;" json:"id"`
	AdminId   string    `gorm:"column:admin_id;type:char(36);not null;default:'';index:idx_admin_id_status,priority:1;comment:业务主键;" json:"admin_id"`
	UserName  string    `gorm:"column:user_name;type:varchar(256);not null;default:'';comment:用户名" json:"user_name"`
	UserPhone string    `gorm:"column:user_phone;type:char(11);not null;default:'';index:idx_phone;comment:手机号" json:"user_phone"`
	Password  string    `gorm:"column:password;type:char(32);not null;default:'';comment:密码" json:"password"`
	Status    string    `gorm:"column:status;type:varchar(32);not null;default:'on';index:idx_admin_id_status,priority:2;comment:状态,off禁用,on启用" json:"status"`
	RoleId    string    `gorm:"column:role_id;type:char(36);not null;default:'';index:role_id;comment:角色Id" json:"role_id"`
	IsInit    string    `gorm:"column:is_init;type:varchar(32);not null;default:'off';comment:是否初始数据,off否,on是" json:"is_init"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;index:idx_created_at;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
}

func (r *CRMAdmin) TableName() string {
	return "crm_admin"
}
