package permission

import (
	"crm/gopkg/gins"
	"crm/handler/api/permission/request"

	"github.com/gin-gonic/gin"
)

// PermissionStatus 管理权限-权限设置-状态
func (h *Handler) PermissionStatus(ctx *gin.Context) {
	var req request.PermissionStatusReq

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.permissionService.PermissionStatus(ctx, req.PermissionId, req.Status)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
