package admin

import (
	"crm/gopkg/gins"
	"crm/internal/service"
	"crm/internal/service/admin"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g            *gin.RouterGroup
	adminService service.AdminIFace
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:            g,
		adminService: admin.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/admin")
	g.POST("/create", h.AdminCreate)
	g.GET("/list", h.AdminList)
	g.POST("/delete", h.AdminDelete)
	g.POST("/edit", h.AdminUpdate)
}
