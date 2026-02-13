package picture_book_category

import (
	"crm/gopkg/gins"
	"crm/handler/api/picture_book_category/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CategoryList(c *gin.Context) {
	var req request.CategoryListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		gins.BadRequest(c, err)
		return
	}
	res, err := h.service.CategoryList(c, req.Offset, req.Limit, req.CategoryName, req.CategoryType)
	if err != nil {
		gins.ServerError(c, err)
		return
	}
	c.JSON(200, res)
}
