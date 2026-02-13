package picture_book

import (
	"crm/gopkg/gins"
	"crm/handler/api/picture_book/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) BookDetail(c *gin.Context) {
	var req request.BookDetailRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		gins.BadRequest(c, err)
		return
	}
	res, err := h.service.BookDetail(c, req.BookId)
	if err != nil {
		gins.ServerError(c, err)
		return
	}
	c.JSON(200, res)
}
