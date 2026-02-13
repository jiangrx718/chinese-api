package api

import (
	"crm/gopkg/gins"
	"crm/handler/api/admin"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	engine *gin.Engine
}

func NewHandler(engine *gin.Engine) gins.Handler {
	return &Handler{
		engine: engine,
	}
}

func (h *Handler) RegisterRoutes() {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	h.engine.Use(cors.New(config))

	// 业务路由
	g := h.engine.Group("/api")
	handlers := []gins.Handler{
		admin.NewHandler(g),
	}

	for _, handler := range handlers {
		handler.RegisterRoutes()
	}
}
