package login

import (
	"crm/gopkg/gins"
	"crm/internal/service"
	"crm/internal/service/login"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g            *gin.RouterGroup
	loginService service.LoginIFace
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:            g,
		loginService: login.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/login")
	g.POST("/do", h.DoLogin)
}
