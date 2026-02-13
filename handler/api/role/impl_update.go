package role

import (
	"crm/gopkg/gins"
	"crm/handler/api/role/request"

	"github.com/gin-gonic/gin"
)

// RoleUpdate 管理权限-角色管理-更新
func (h *Handler) RoleUpdate(ctx *gin.Context) {
	var req request.RoleUpdateReq

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.roleService.RoleUpdate(ctx, req.RoleId, req.RoleName, req.Status, req.Permission)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
