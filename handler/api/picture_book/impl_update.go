package picture_book

import (
	"crm/gopkg/gins"
	"crm/handler/api/picture_book/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) BookUpdate(c *gin.Context) {
	var req request.BookUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		gins.BadRequest(c, err)
		return
	}
	res, err := h.service.BookUpdate(c, req.BookId, req.CategoryId, req.Title, req.Icon, req.Status, req.Position, req.CType)
	if err != nil {
		gins.ServerError(c, err)
		return
	}
	c.JSON(200, res)
}
