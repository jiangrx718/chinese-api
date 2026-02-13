package permission

import (
	"crm/gopkg/utils/httputil"
	"crm/handler/api/permission/request"

	"github.com/gin-gonic/gin"
)

// PermissionMenu 管理权限-左侧菜单-列表
func (h *Handler) PermissionMenu(ctx *gin.Context) {

	var query request.ListQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		httputil.BadRequest(ctx, err)
		return
	}

	loginUserId, _ := ctx.Get("user_id")

	result, err := h.permissionService.PermissionMenu(ctx, loginUserId.(string))
	if err != nil {
		httputil.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
