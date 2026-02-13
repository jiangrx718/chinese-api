package admin

import (
	"crm/gopkg/gins"
	"crm/handler/api/admin/request"

	"github.com/gin-gonic/gin"
)

// AdminCreate 管理权限-管理员-创建
func (h *Handler) AdminCreate(ctx *gin.Context) {
	var req request.AdminCreateReq

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.adminService.AdminCreate(ctx, req.UserName, req.UserPhone, req.Password, req.Status, req.RoleId)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
