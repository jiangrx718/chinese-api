package picture_book_item

import (
	"crm/gopkg/gins"
	"crm/handler/api/picture_book_item/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ItemDelete(c *gin.Context) {
	var req request.ItemDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		gins.BadRequest(c, err)
		return
	}
	res, err := h.service.ItemDelete(c, req.Id)
	if err != nil {
		gins.ServerError(c, err)
		return
	}
	c.JSON(200, res)
}
