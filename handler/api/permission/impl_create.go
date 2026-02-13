package permission

import (
	"crm/gopkg/gins"
	"crm/handler/api/permission/request"

	"github.com/gin-gonic/gin"
)

// PermissionCreate 管理权限-权限设置-创建
func (h *Handler) PermissionCreate(ctx *gin.Context) {
	var req request.PermissionCreateReq

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.permissionService.PermissionCreate(ctx, req.PermissionName, req.PermissionUrl, req.ParentId, req.Status, req.Position, req.PermissionType)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
