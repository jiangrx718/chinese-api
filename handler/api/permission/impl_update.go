package permission

import (
	"crm/gopkg/gins"
	"crm/handler/api/permission/request"

	"github.com/gin-gonic/gin"
)

// PermissionUpdate 管理权限-权限设置-更新
func (h *Handler) PermissionUpdate(ctx *gin.Context) {
	var req request.PermissionUpdateReq

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.permissionService.PermissionUpdate(ctx, req.PermissionId, req.PermissionName, req.PermissionUrl, req.ParentId, req.Status, req.Position)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
