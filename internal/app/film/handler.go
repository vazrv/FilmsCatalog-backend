// internal/app/film/handler.go
package film

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetPopular(c *gin.Context) {
	films, err := h.service.GetPopularFilms(10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to load films"})
		return
	}

	// Отправляем напрямую — GORM уже положил валидный JSON в `genres`
	c.JSON(http.StatusOK, films)
}