package permission

import (
	"crm/gopkg/gins"
	"crm/internal/service"
	"crm/internal/service/permission"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g                 *gin.RouterGroup
	permissionService service.PermissionIFace
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:                 g,
		permissionService: permission.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/permission")
	g.POST("/create", h.PermissionCreate)
	g.GET("/list", h.PermissionList)
	g.GET("/menu", h.PermissionMenu)
	g.POST("/delete", h.PermissionDelete)
	g.POST("/edit", h.PermissionUpdate)
	g.POST("/status", h.PermissionStatus)
}
