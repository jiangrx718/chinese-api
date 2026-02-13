package role

import (
	"crm/gopkg/gins"
	"crm/internal/service"
	"crm/internal/service/role"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g           *gin.RouterGroup
	roleService service.RoleIFace
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:           g,
		roleService: role.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/role")
	g.POST("/create", h.RoleCreate)
	g.GET("/list", h.RoleList)
	g.POST("/delete", h.RoleDelete)
	g.POST("/edit", h.RoleUpdate)
}
