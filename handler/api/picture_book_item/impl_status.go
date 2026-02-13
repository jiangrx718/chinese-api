package picture_book_item

import (
	"crm/gopkg/gins"
	"crm/handler/api/picture_book_item/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ItemStatus(c *gin.Context) {
	var req request.ItemStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		gins.BadRequest(c, err)
		return
	}
	res, err := h.service.ItemStatus(c, req.Id, req.Status)
	if err != nil {
		gins.ServerError(c, err)
		return
	}
	c.JSON(200, res)
}
