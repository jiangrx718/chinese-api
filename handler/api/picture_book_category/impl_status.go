package picture_book_category

import (
	"crm/gopkg/gins"
	"crm/handler/api/picture_book_category/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CategoryStatus(c *gin.Context) {
	var req request.CategoryStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		gins.BadRequest(c, err)
		return
	}
	res, err := h.service.CategoryStatus(c, req.CategoryId, req.Status)
	if err != nil {
		gins.ServerError(c, err)
		return
	}
	c.JSON(200, res)
}
