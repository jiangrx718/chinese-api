package category

import (
	"crm/gopkg/gins"
	"crm/internal/service"
	"crm/internal/service/category"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g               *gin.RouterGroup
	categoryService service.CategoryIFace
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:               g,
		categoryService: category.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/book")
	g.GET("/category", h.CategoryList)
}
