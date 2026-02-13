package file

import (
	"crm/gopkg/gins"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g *gin.RouterGroup
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g: g,
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/file")
	g.POST("/upload", h.FileUpload)
}
