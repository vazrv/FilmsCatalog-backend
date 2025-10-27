// internal/app/actor/handler.go
package actor

import "github.com/gin-gonic/gin"

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetTop(c *gin.Context) {
	actors, err := h.service.GetTopActors(10)
	if err != nil {
		c.JSON(500, gin.H{"message": "Failed to load actors"})
		return
	}
	c.JSON(200, actors)
}