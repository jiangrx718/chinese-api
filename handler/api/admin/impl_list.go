package admin

import (
	"crm/gopkg/utils/httputil"
	"crm/handler/api/admin/request"

	"github.com/gin-gonic/gin"
)

// AdminList 管理权限-管理员-列表
func (h *Handler) AdminList(ctx *gin.Context) {

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

	result, err := h.adminService.AdminList(ctx, query.Offset, query.Limit, query.Status, query.UserPhone)
	if err != nil {
		httputil.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
