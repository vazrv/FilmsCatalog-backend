package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetProfile — возвращает профиль текущего пользователя
func (h *UserHandler) GetProfile(c *gin.Context) {
	// Получаем userID из middleware (JWT)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Пользователь не авторизован"})
		return
	}

	profile, err := h.service.GetProfile(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Ошибка при получении профиля"})
		return
	}

	c.JSON(http.StatusOK, profile)
}
