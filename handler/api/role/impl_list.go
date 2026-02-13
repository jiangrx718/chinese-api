package role

import (
	"crm/gopkg/utils/httputil"
	"crm/handler/api/role/request"

	"github.com/gin-gonic/gin"
)

// RoleList  管理权限-角色管理-列表
func (h *Handler) RoleList(ctx *gin.Context) {

	var query request.ListQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		httputil.BadRequest(ctx, err)
		return
	}
	if query.Offset >= 1 {
		query.Offset -= 1
		query.Offset *= query.Limit
	}
	if query.Limit > request.MaxLimit {
		query.Limit = request.MaxLimit
	}

	result, err := h.roleService.RoleList(ctx, query.Offset, query.Limit)
	if err != nil {
		httputil.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
