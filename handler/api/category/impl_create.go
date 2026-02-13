package category

import (
	"crm/gopkg/gins"
	"crm/handler/api/category/request"

	"github.com/gin-gonic/gin"
)

// CategoryCreate 内容管理-文章/商品分类-创建
func (h *Handler) CategoryCreate(ctx *gin.Context) {
	var req request.CategoryCreateReq

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.categoryService.CategoryCreate(ctx, req.ParentId, req.CategoryName, req.CategoryImage, req.Status, req.CategoryType, req.Position)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
