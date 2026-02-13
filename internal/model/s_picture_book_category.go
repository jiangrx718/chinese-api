package model

import "time"

// SPictureBookCategory 绘本类型表
type SPictureBookCategory struct {
	Id           int       `gorm:"column:id;type:bigint;not null;primaryKey;autoIncrement;comment:主键;" json:"id"`
	CategoryId   string    `gorm:"column:category_id;type:char(36);not null;default:'';index:idx_category_id;;comment:栏目ID;" json:"category_id"`
	CategoryName string    `gorm:"column:category_name;type:varchar(256);not null;default:'';comment:栏目名称" json:"category_name"`
	Position     int       `gorm:"column:position;type:int;not null;default:0;comment:排序,倒序" json:"position"`
	Type         int       `gorm:"column:type;type:int;not null;default:0;comment:1中文绘本,2英文绘本,3古诗绘本,4英语词汇" json:"type"`
	Status       string    `gorm:"column:status;type:varchar(20);not null;default:'on';comment:状态,on启用,off禁用" json:"status"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;index:idx_created_at;comment:创建时间" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
}

func (r *SPictureBookCategory) TableName() string {
	return "s_picture_book_category"
}
