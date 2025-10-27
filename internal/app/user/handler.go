// internal/app/user/handler.go
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

func (h *UserHandler) GetProfile(c *gin.Context) {
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

	// Формируем ответ в нужном формате
	c.JSON(http.StatusOK, gin.H{
		"id":           profile.ID,
		"username":     profile.Username,
		"email":        profile.Email,
		"avatar_url":   profile.AvatarURL,
		"created_at":   profile.CreatedAt,
		"is_admin":     profile.IsAdmin,
		"stats": gin.H{
			"favorites_count": profile.FavoritesCount,
			"reviews_count":   profile.ReviewsCount,
			"ratings_count":   profile.RatingsCount,
		},
	})
}