package request

import "crm/gopkg/utils/httputil"

// AdminCreateReq 创建参数
type AdminCreateReq struct {
	UserName  string `json:"user_name" binding:"required"`
	UserPhone string `json:"user_phone" binding:"required"`
	Password  string `json:"password" binding:"required"`
	RoleId    string `json:"role_id" binding:"required"`
	Status    string `json:"status" binding:"required"`
}

// AdminDeleteReq 删除参数
type AdminDeleteReq struct {
	AdminId string `json:"admin_id" binding:"required"`
}

// AdminUpdateReq 修改参数
type AdminUpdateReq struct {
	AdminId  string `json:"admin_id" binding:"required"`
	Password string `json:"password"`
	RoleId   string `json:"role_id" binding:"required"`
	Status   string `json:"status" binding:"required"`
}

// ListQuery 列表参数
type ListQuery struct {
	httputil.Pagination
	Status    string `json:"status" form:"status"`
	UserPhone string `json:"user_phone" form:"user_phone"`
}

const (
	MaxLimit = 100
)
