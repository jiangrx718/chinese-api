package article

import (
	"crm/gopkg/utils/httputil"
	"crm/handler/api/article/request"

	"github.com/gin-gonic/gin"
)

// ArticleDetail 内容管理-文章列表-详情
func (h *Handler) ArticleDetail(ctx *gin.Context) {
	var query request.ArticleDetailReq
	if err := ctx.ShouldBindQuery(&query); err != nil {
		httputil.BadRequest(ctx, err)
		return
	}
	result, err := h.articleService.ArticleDetail(ctx, query.ArticleId)
	if err != nil {
		httputil.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
