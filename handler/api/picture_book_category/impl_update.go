package picture_book_category

import (
	"crm/gopkg/gins"
	"crm/handler/api/picture_book_category/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CategoryUpdate(c *gin.Context) {
	var req request.CategoryUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		gins.BadRequest(c, err)
		return
	}
	res, err := h.service.CategoryUpdate(c, req.CategoryId, req.CategoryName, req.Status, req.Position, req.CategoryType)
	if err != nil {
		gins.ServerError(c, err)
		return
	}
	c.JSON(200, res)
}
