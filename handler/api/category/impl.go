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
	p := h.g.Group("/category/product")
	p.POST("/create", h.CategoryCreate)
	p.POST("/update", h.CategoryUpdate)
	p.POST("/delete", h.CategoryDelete)
	p.POST("/status", h.CategoryStatus)
	p.GET("/list", h.CategoryList)

	a := h.g.Group("/category/article")
	a.POST("/create", h.CategoryCreate)
	a.POST("/update", h.CategoryUpdate)
	a.POST("/delete", h.CategoryDelete)
	a.POST("/status", h.CategoryStatus)
	a.GET("/list", h.CategoryList)
}
