package category

import (
	"crm/gopkg/utils/httputil"

	"github.com/gin-gonic/gin"
)

// CategoryListQuery 列表参数
type CategoryListQuery struct {
	Type string `json:"type" form:"type"`
}

func (h *Handler) CategoryList(ctx *gin.Context) {
	var query CategoryListQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		httputil.BadRequest(ctx, err)
		return
	}

	result, err := h.categoryService.CategoryList(ctx, query.Type)
	if err != nil {
		httputil.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
}
