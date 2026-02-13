package book

import (
	"crm/gopkg/gins"
	"crm/internal/service"
	"crm/internal/service/book"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g           *gin.RouterGroup
	bookService service.BookIFace
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:           g,
		bookService: book.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/book")
	g.GET("/list", h.BookList)
}
