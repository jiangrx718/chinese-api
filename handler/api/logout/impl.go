package logout

import (
	"crm/gopkg/gins"
	"crm/internal/service"
	"crm/internal/service/logout"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g             *gin.RouterGroup
	logoutService service.LogoutIFace
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:             g,
		logoutService: logout.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/logout")
	g.POST("/do", h.DoLogout)
}
