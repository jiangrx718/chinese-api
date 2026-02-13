package picture_book_item

import (
	"crm/gopkg/gins"
	"crm/handler/api/picture_book_item/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ItemUpdate(c *gin.Context) {
	var req request.ItemUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		gins.BadRequest(c, err)
		return
	}
	res, err := h.service.ItemUpdate(c, req.Id, req.Title, req.Pic, req.BPic, req.Audio, req.Content, req.Status, req.Position)
	if err != nil {
		gins.ServerError(c, err)
		return
	}
	c.JSON(200, res)
}
