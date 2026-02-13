package permission

import (
	"crm/gopkg/gins"
	"crm/handler/api/permission/request"

	"github.com/gin-gonic/gin"
)

// PermissionDelete 管理权限-权限设置-删除
func (h *Handler) PermissionDelete(ctx *gin.Context) {
	var req request.PermissionDeleteReq

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.permissionService.PermissionDelete(ctx, req.PermissionId)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
