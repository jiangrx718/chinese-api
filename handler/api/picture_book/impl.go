package picture_book

import (
	"crm/gopkg/gins"
	"crm/internal/service"
	"crm/internal/service/picture_book"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g       *gin.RouterGroup
	service service.PictureBookIFace
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:       g,
		service: picture_book.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/picture/book")
	g.POST("/create", h.BookCreate)
	g.POST("/update", h.BookUpdate)
	g.POST("/delete", h.BookDelete)
	g.POST("/status", h.BookStatus)
	g.GET("/list", h.BookList)
	g.GET("/detail", h.BookDetail)
}
