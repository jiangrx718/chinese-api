package category

import (
	"crm/gopkg/gins"
	"crm/handler/api/category/request"

	"github.com/gin-gonic/gin"
)

// CategoryUpdate 内容管理-文章/商品分类-更新
func (h *Handler) CategoryUpdate(ctx *gin.Context) {
	var req request.CategoryUpdateReq

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.categoryService.CategoryUpdate(ctx, req.CategoryId, req.ParentId, req.CategoryName, req.CategoryImage, req.Status, req.CategoryType, req.Position)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
