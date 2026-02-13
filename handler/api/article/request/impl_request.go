package request

import "crm/gopkg/utils/httputil"

// ArticleCreateReq 创建参数
type ArticleCreateReq struct {
	CategoryId     string `json:"category_id" binding:"required"`
	ArticleName    string `json:"article_name" binding:"required"`
	ArticleImage   string `json:"article_image"`
	Position       int    `json:"position"`
	Status         string `json:"status" binding:"required"`
	ArticleContent string `json:"article_content" binding:"required"`
}

// ArticleUpdateReq 修改参数
type ArticleUpdateReq struct {
	ArticleId      string `json:"article_id" binding:"required"`
	CategoryId     string `json:"category_id" binding:"required"`
	ArticleName    string `json:"article_name" binding:"required"`
	ArticleImage   string `json:"article_image"`
	Position       int    `json:"position"`
	Status         string `json:"status" binding:"required"`
	ArticleContent string `json:"article_content" binding:"required"`
}

// ArticleDeleteReq 删除参数
type ArticleDeleteReq struct {
	ArticleId string `json:"article_id" binding:"required"`
}

// ArticleStatusReq 状态参数
type ArticleStatusReq struct {
	ArticleId string `json:"article_id" binding:"required"`
	Status    string `json:"status" binding:"required"`
}

// ListQuery 列表参数
type ListQuery struct {
	httputil.Pagination
	CategoryId  string `json:"category_id" form:"category_id"`
	ArticleName string `json:"article_name" form:"article_name"`
}

const (
	MaxLimit = 100
)

// ArticleDetailReq 详情参数
type ArticleDetailReq struct {
	ArticleId string `json:"article_id" form:"article_id"`
}
