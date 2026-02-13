package book

import (
	"crm/gopkg/utils/httputil"

	"github.com/gin-gonic/gin"
)

// BookListQuery 列表参数
type BookListQuery struct {
	Type int `json:"type" form:"type"`
}

func (h *Handler) BookList(ctx *gin.Context) {
	var query BookListQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		httputil.BadRequest(ctx, err)
		return
	}

	result, err := h.bookService.BookList(ctx, query.Type)
	if err != nil {
		httputil.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
