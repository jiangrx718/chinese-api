package info

import (
	"crm/gopkg/utils/httputil"

	"github.com/gin-gonic/gin"
)

// InfoListQuery 列表参数
type InfoListQuery struct {
	BookId string `json:"book_id" form:"book_id"`
}

func (h *Handler) InfoList(ctx *gin.Context) {
	var query InfoListQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		httputil.BadRequest(ctx, err)
		return
	}

	result, err := h.infoService.InfoList(ctx, query.BookId)
	if err != nil {
		httputil.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
