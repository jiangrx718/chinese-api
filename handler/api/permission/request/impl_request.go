package request

// PermissionCreateReq 创建参数
type PermissionCreateReq struct {
	PermissionName string `json:"permission_name" binding:"required"`
	PermissionUrl  string `json:"permission_url"`
	ParentId       string `json:"parent_id"`
	Position       int    `json:"position"`
	PermissionType int    `json:"permission_type" binding:"required"`
	Status         string `json:"status" binding:"required"`
}

// PermissionDeleteReq 删除参数
type PermissionDeleteReq struct {
	PermissionId string `json:"permission_id" binding:"required"`
}

// PermissionUpdateReq 修改参数
type PermissionUpdateReq struct {
	PermissionId   string `json:"permission_id" binding:"required"`
	PermissionName string `json:"permission_name" binding:"required"`
	PermissionUrl  string `json:"permission_url"`
	ParentId       string `json:"parent_id"`
	Position       int    `json:"position"`
	Status         string `json:"status" binding:"required"`
}

// PermissionStatusReq 状态参数
type PermissionStatusReq struct {
	PermissionId string `json:"permission_id" binding:"required"`
	Status       string `json:"status" binding:"required"`
}

// ListQuery 列表参数
type ListQuery struct {
	Status string `json:"status" form:"status"`
}
