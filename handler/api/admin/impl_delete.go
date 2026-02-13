package admin

import (
	"crm/gopkg/gins"
	"crm/handler/api/admin/request"

	"github.com/gin-gonic/gin"
)

// AdminDelete 管理权限-管理员-删除
func (h *Handler) AdminDelete(ctx *gin.Context) {
	var req request.AdminDeleteReq

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.adminService.AdminDelete(ctx, req.AdminId)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
