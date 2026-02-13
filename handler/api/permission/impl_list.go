package permission

import (
	"crm/gopkg/utils/httputil"
	"crm/handler/api/permission/request"

	"github.com/gin-gonic/gin"
)

// PermissionList 管理权限-权限设置-列表
func (h *Handler) PermissionList(ctx *gin.Context) {

	var query request.ListQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		httputil.BadRequest(ctx, err)
		return
	}
	result, err := h.permissionService.PermissionList(ctx, query.Status)
	if err != nil {
		httputil.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
