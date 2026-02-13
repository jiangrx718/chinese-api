package article

import (
	"crm/gopkg/gins"
	"crm/handler/api/article/request"

	"github.com/gin-gonic/gin"
)

// ArticleCreate 内容管理-文章列表-创建
func (h *Handler) ArticleCreate(ctx *gin.Context) {
	var req request.ArticleCreateReq

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.articleService.ArticleCreate(ctx, req.CategoryId, req.ArticleName, req.ArticleImage, req.Status, req.ArticleContent, req.Position)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
