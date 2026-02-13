package model

import "time"

// CRMArticle 文章表
type CRMArticle struct {
	Id           int       `gorm:"column:id;type:bigint;not null;primaryKey;autoIncrement;comment:主键;" json:"id"`
	ArticleId    string    `gorm:"column:article_id;type:char(36);not null;default:'';comment:业务主键;" json:"article_id"`
	CategoryId   string    `gorm:"column:category_id;type:char(36);not null;default:'';index:idx_category_id;comment:栏目Id;" json:"category_id"`
	ArticleName  string    `gorm:"column:article_name;type:varchar(1024);not null;default:'';comment:文章名称" json:"article_name"`
	ArticleImage string    `gorm:"column:article_image;type:varchar(1024);not null;default:'';comment:文章缩略图" json:"article_image"`
	Status       string    `gorm:"column:status;type:varchar(32);not null;default:'on';comment:状态,off禁用,on启用" json:"status"`
	Position     int       `gorm:"column:position;type:int;not null;default:0;comment:排序" json:"position"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;index:idx_created_at;comment:创建时间" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
}

func (r *CRMArticle) TableName() string {
	return "crm_article"
}

// CRMArticleContent 文章内容表
type CRMArticleContent struct {
	Id             int       `gorm:"column:id;type:bigint;not null;primaryKey;autoIncrement;comment:主键;" json:"id"`
	ContentId      string    `gorm:"column:content_id;type:char(36);not null;default:'';comment:业务主键;" json:"content_id"`
	ArticleId      string    `gorm:"column:article_id;type:char(36);not null;default:'';index:idx_article_id;comment:文章Id;" json:"article_id"`
	ArticleContent string    `gorm:"column:article_content;type:text;comment:文章内容" json:"article_content"`
	CreatedAt      time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;index:idx_created_at;comment:创建时间" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
}

func (r *CRMArticleContent) TableName() string {
	return "crm_article_content"
}
