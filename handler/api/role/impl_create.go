package role

import (
	"crm/gopkg/gins"
	"crm/handler/api/role/request"

	"github.com/gin-gonic/gin"
)

// RoleCreate 管理权限-角色管理-创建
func (h *Handler) RoleCreate(ctx *gin.Context) {
	var req request.RoleCreateReq

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.roleService.RoleCreate(ctx, req.RoleName, req.Status, req.Permission)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
