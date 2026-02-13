package picture

import (
	"crm/gopkg/utils/httputil"

	"github.com/gin-gonic/gin"
)

// PictureListQuery 列表参数
type PictureListQuery struct {
	Type int `json:"type" form:"type"`
	httputil.Pagination
}

func (h *Handler) PictureList(ctx *gin.Context) {
	var query PictureListQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		httputil.BadRequest(ctx, err)
		return
	}

	result, err := h.pictureService.PictureList(ctx, query.Type, query.Offset, query.Limit)
	if err != nil {
		httputil.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
