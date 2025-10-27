// internal/app/genre/handler.go
package genre

import "github.com/gin-gonic/gin"

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAll(c *gin.Context) {
	genres, err := h.service.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"message": "Failed to load genres"})
		return
	}
	c.JSON(200, genres)
}