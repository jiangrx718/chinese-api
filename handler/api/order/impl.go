package order

import (
	"crm/gopkg/gins"
	"crm/internal/service"
	"crm/internal/service/order"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g            *gin.RouterGroup
	orderService service.OrderIFace
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:            g,
		orderService: order.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	//g := h.g.Group("/admin")
	//g.POST("/create", h.CreateData)
	//g.POST("/update", h.UpdateData)
	//g.POST("/delete", h.DeleteData)
	//g.POST("/status", h.StatusData)
	//g.GET("/list", h.ListData)
}
