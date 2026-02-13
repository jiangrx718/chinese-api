package article

import (
	"crm/gopkg/gins"
	"crm/handler/api/article/request"

	"github.com/gin-gonic/gin"
)

// ArticleDelete 内容管理-文章列表-删除
func (h *Handler) ArticleDelete(ctx *gin.Context) {
	var req request.ArticleDeleteReq

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.articleService.ArticleDelete(ctx, req.ArticleId)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
