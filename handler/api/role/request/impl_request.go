package request

import "crm/gopkg/utils/httputil"

// RoleCreateReq 创建参数
type RoleCreateReq struct {
	RoleName   string   `json:"role_name" binding:"required"`
	Status     string   `json:"status" binding:"required"`
	Permission []string `json:"permission"`
}

// RoleDeleteReq 删除参数
type RoleDeleteReq struct {
	RoleId string `json:"role_id" binding:"required"`
}

// RoleUpdateReq 修改参数
type RoleUpdateReq struct {
	RoleId     string   `json:"role_id" binding:"required"`
	RoleName   string   `json:"role_name" binding:"required"`
	Status     string   `json:"status" binding:"required"`
	Permission []string `json:"permission"`
}

// ListQuery 列表参数
type ListQuery struct {
	httputil.Pagination
}

const (
	MaxLimit = 100
)
