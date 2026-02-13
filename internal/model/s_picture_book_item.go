package model

import "time"

// SPictureBookItem 绘本详情
type SPictureBookItem struct {
	Id        int       `gorm:"column:id;type:bigint;not null;primaryKey;autoIncrement;comment:主键;" json:"id"`
	BookId    string    `gorm:"column:book_id;type:char(36);not null;default:'';index:idx_book_id;;comment:绘本id;" json:"book_id"`
	Title     string    `gorm:"column:title;type:varchar(256);not null;default:'';comment:绘本标题" json:"title"`
	Pic       string    `gorm:"column:pic;type:varchar(1024);not null;default:'';comment:绘本详情图" json:"pic"`
	BPic      string    `gorm:"column:b_pic;type:varchar(1024);not null;default:'';comment:绘本详情大图" json:"b_pic"`
	Audio     string    `gorm:"column:audio;type:varchar(1024);not null;default:'';comment:绘本详音频" json:"audio"`
	Status    string    `gorm:"column:status;type:varchar(20);not null;default:'on';comment:状态,on启用,off禁用" json:"status"`
	Content   string    `gorm:"column:content;type:varchar(4096);not null;default:'';comment:原文" json:"content"`
	Position  int       `gorm:"column:position;type:int;not null;default:0;comment:排序,倒序" json:"position"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;index:idx_created_at;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
}

func (r *SPictureBookItem) TableName() string {
	return "s_picture_book_item"
}
