package article

import (
	"crm/gopkg/gins"
	"crm/internal/service"
	"crm/internal/service/article"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g              *gin.RouterGroup
	articleService service.ArticleIFace
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:              g,
		articleService: article.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/article")
	g.POST("/create", h.ArticleCreate)
	g.POST("/update", h.ArticleUpdate)
	g.POST("/delete", h.ArticleDelete)
	g.POST("/status", h.ArticleStatus)
	g.GET("/list", h.ArticleList)
	g.GET("/detail", h.ArticleDetail)
}
