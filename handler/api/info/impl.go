package info

import (
	"crm/gopkg/gins"
	"crm/handler/middleware"
	"crm/internal/service"
	"crm/internal/service/info"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g           *gin.RouterGroup
	infoService service.InfoIFace
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:           g,
		infoService: info.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/info")
	g.GET("/list", middleware.Signature(), h.InfoList)
}
