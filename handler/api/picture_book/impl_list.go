package picture_book

import (
	"crm/gopkg/gins"
	"crm/handler/api/picture_book/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) BookList(c *gin.Context) {
	var req request.BookListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		gins.BadRequest(c, err)
		return
	}
	res, err := h.service.BookList(c, req.Offset, req.Limit, req.Title, req.CType, req.CategoryId)
	if err != nil {
		gins.ServerError(c, err)
		return
	}
	c.JSON(200, res)
}
