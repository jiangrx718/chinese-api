package article

import (
	"crm/gopkg/gins"
	"crm/handler/api/article/request"

	"github.com/gin-gonic/gin"
)

// ArticleStatus 内容管理-文章列表-状态
func (h *Handler) ArticleStatus(ctx *gin.Context) {
	var req request.ArticleStatusReq

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.articleService.ArticleStatus(ctx, req.ArticleId, req.Status)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
