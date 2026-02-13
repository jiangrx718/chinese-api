package request

import "crm/gopkg/utils/httputil"

// CategoryCreateReq 创建参数
type CategoryCreateReq struct {
	ParentId      string `json:"parent_id"`
	CategoryImage string `json:"category_image"`
	CategoryType  int    `json:"category_type" binding:"required"`
	Position      int    `json:"position"`
	CategoryName  string `json:"category_name" binding:"required"`
	Status        string `json:"status" binding:"required"`
}

// CategoryUpdateReq 修改参数
type CategoryUpdateReq struct {
	CategoryId    string `json:"category_id" binding:"required"`
	ParentId      string `json:"parent_id"`
	CategoryImage string `json:"category_image"`
	CategoryType  int    `json:"category_type" binding:"required"`
	Position      int    `json:"position"`
	CategoryName  string `json:"category_name" binding:"required"`
	Status        string `json:"status" binding:"required"`
}

// CategoryDeleteReq 删除参数
type CategoryDeleteReq struct {
	CategoryId string `json:"category_id" binding:"required"`
}

// CategoryStatusReq 状态参数
type CategoryStatusReq struct {
	CategoryId string `json:"category_id" binding:"required"`
	Status     string `json:"status" binding:"required"`
}

// ListQuery 列表参数
type ListQuery struct {
	httputil.Pagination
	CategoryType int `json:"category_type" form:"category_type"`
}

const (
	MaxLimit = 100
)
