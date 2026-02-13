package model

import "time"

// CRMCategory 栏目表
type CRMCategory struct {
	Id            int       `gorm:"column:id;type:bigint;not null;primaryKey;autoIncrement;comment:主键;" json:"id"`
	CategoryId    string    `gorm:"column:category_id;type:char(36);not null;default:'';comment:业务主键;" json:"category_id"`
	CategoryName  string    `gorm:"column:category_name;type:varchar(128);not null;default:'';comment:栏目名称" json:"category_name"`
	CategoryImage string    `gorm:"column:category_image;type:varchar(1024);not null;default:'';comment:栏目图片" json:"category_image"`
	CategoryType  int       `gorm:"column:category_type;type:int;not null;default:0;comment:栏目类型,1文章,2商品" json:"category_type"`
	ParentId      string    `gorm:"column:parent_id;type:char(36);not null;default:'';comment:父级id;" json:"parent_id"`
	Status        string    `gorm:"column:status;type:varchar(32);not null;default:'on';comment:状态,off禁用,on启用" json:"status"`
	Position      int       `gorm:"column:position;type:int;not null;default:0;comment:排序" json:"position"`
	CreatedAt     time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;index:idx_created_at;comment:创建时间" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
}

func (r *CRMCategory) TableName() string {
	return "crm_category"
}
