package picture_book_item

import (
	"crm/gopkg/gins"
	"crm/internal/service"
	"crm/internal/service/picture_book_item"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g       *gin.RouterGroup
	service service.PictureBookItemIFace
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:       g,
		service: picture_book_item.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/picture/book/item")
	g.POST("/create", h.ItemCreate)
	g.POST("/update", h.ItemUpdate)
	g.POST("/delete", h.ItemDelete)
	g.POST("/status", h.ItemStatus)
	g.GET("/:book_id/list", h.ItemList)
}
