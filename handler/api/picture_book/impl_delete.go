package picture_book

import (
	"crm/gopkg/gins"
	"crm/handler/api/picture_book/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) BookDelete(c *gin.Context) {
	var req request.BookDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		gins.BadRequest(c, err)
		return
	}
	res, err := h.service.BookDelete(c, req.BookId)
	if err != nil {
		gins.ServerError(c, err)
		return
	}
	c.JSON(200, res)
}
