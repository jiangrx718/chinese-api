package api

import (
	"crm/gopkg/gins"
	"crm/handler/api/admin"
	"crm/handler/api/article"
	"crm/handler/api/category"
	"crm/handler/api/file"
	"crm/handler/api/login"
	"crm/handler/api/logout"
	"crm/handler/api/order"
	"crm/handler/api/permission"
	"crm/handler/api/picture_book"
	"crm/handler/api/picture_book_category"
	"crm/handler/api/picture_book_item"
	"crm/handler/api/product"
	"crm/handler/api/role"
	"crm/handler/api/system"
	"crm/handler/middleware"

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

	// 登录路由
	l := h.engine.Group("/")
	loginHandlers := []gins.Handler{
		login.NewHandler(l),
	}
	for _, handler := range loginHandlers {
		handler.RegisterRoutes()
	}

	// 退出路由
	lg := h.engine.Group("/api", middleware.JWTAuth())
	logoutHandlers := []gins.Handler{
		logout.NewHandler(lg),
	}
	for _, handler := range logoutHandlers {
		handler.RegisterRoutes()
	}

	// 业务路由
	g := h.engine.Group("/api", middleware.JWTAuth(), middleware.PermissionAuth())
	handlers := []gins.Handler{
		admin.NewHandler(g),
		article.NewHandler(g),
		category.NewHandler(g),
		file.NewHandler(g),
		order.NewHandler(g),
		permission.NewHandler(g),
		picture_book.NewHandler(g),
		picture_book_category.NewHandler(g),
		picture_book_item.NewHandler(g),
		product.NewHandler(g),
		role.NewHandler(g),
		system.NewHandler(g),
	}

	for _, handler := range handlers {
		handler.RegisterRoutes()
	}
}
