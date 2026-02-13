package picture_book_category

import (
	"crm/gopkg/gins"
	"crm/handler/api/picture_book_category/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CategoryDelete(c *gin.Context) {
	var req request.CategoryDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		gins.BadRequest(c, err)
		return
	}
	res, err := h.service.CategoryDelete(c, req.CategoryId)
	if err != nil {
		gins.ServerError(c, err)
		return
	}
	c.JSON(200, res)
}
