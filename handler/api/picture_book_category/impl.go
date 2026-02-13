package picture_book_category

import (
	"crm/gopkg/gins"
	"crm/internal/service"
	"crm/internal/service/picture_book_category"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g       *gin.RouterGroup
	service service.PictureBookCategoryIFace
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:       g,
		service: picture_book_category.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/picture/book/category")
	g.POST("/create", h.CategoryCreate)
	g.POST("/update", h.CategoryUpdate)
	g.POST("/delete", h.CategoryDelete)
	g.POST("/status", h.CategoryStatus)
	g.GET("/list", h.CategoryList)
}
