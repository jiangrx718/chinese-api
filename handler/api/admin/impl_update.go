package admin

import (
	"crm/gopkg/gins"
	"crm/handler/api/admin/request"

	"github.com/gin-gonic/gin"
)

// AdminUpdate 管理权限-管理员-更新
func (h *Handler) AdminUpdate(ctx *gin.Context) {
	var req request.AdminUpdateReq

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.adminService.AdminUpdate(ctx, req.AdminId, req.Password, req.Status, req.RoleId)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
