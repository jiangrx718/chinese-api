package category

import (
	"crm/gopkg/utils/httputil"
	"crm/handler/api/category/request"
	"github.com/gin-gonic/gin"
)

// CategoryList 内容管理-文章/商品分类-列表
func (h *Handler) CategoryList(ctx *gin.Context) {

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

	result, err := h.categoryService.CategoryList(ctx, query.Offset, query.Limit, query.CategoryType)
	if err != nil {
		httputil.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
