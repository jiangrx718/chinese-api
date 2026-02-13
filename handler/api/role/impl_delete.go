package role

import (
	"crm/gopkg/gins"
	"crm/handler/api/role/request"

	"github.com/gin-gonic/gin"
)

// RoleDelete 管理权限-角色管理-删除
func (h *Handler) RoleDelete(ctx *gin.Context) {
	var req request.RoleDeleteReq

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.roleService.RoleDelete(ctx, req.RoleId)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
