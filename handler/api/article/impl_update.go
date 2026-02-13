package article

import (
	"crm/gopkg/gins"
	"crm/handler/api/article/request"

	"github.com/gin-gonic/gin"
)

// ArticleUpdate 内容管理-文章列表-更新
func (h *Handler) ArticleUpdate(ctx *gin.Context) {
	var req request.ArticleUpdateReq

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.articleService.ArticleUpdate(ctx, req.ArticleId, req.CategoryId, req.ArticleName, req.ArticleImage, req.Status, req.ArticleContent, req.Position)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
