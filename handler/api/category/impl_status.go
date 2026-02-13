package category

import (
	"crm/gopkg/gins"
	"crm/handler/api/category/request"

	"github.com/gin-gonic/gin"
)

// CategoryStatus 内容管理-文章/商品分类-状态
func (h *Handler) CategoryStatus(ctx *gin.Context) {
	var req request.CategoryStatusReq

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.categoryService.CategoryStatus(ctx, req.CategoryId, req.Status)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
