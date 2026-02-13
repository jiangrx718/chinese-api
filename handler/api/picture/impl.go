package picture

import (
	"crm/gopkg/gins"
	"crm/internal/service"
	"crm/internal/service/picture"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g              *gin.RouterGroup
	pictureService service.PictureIFace
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:              g,
		pictureService: picture.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/picture")
	g.GET("/list", h.PictureList)
}
