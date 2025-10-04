// internal/middleware/JWTAuth.go
// Middleware для проверки JWT токена авторизации.
package middleware

import (
	"net/http"
	"strings"

	"FilmsCatalog/internal/config"
	"FilmsCatalog/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTAuth — middleware, который проверяет наличие и валидность JWT токена
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.Get()

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Требуется авторизация"})
			c.Abort()
			return
		}

		// Проверяем формат: Bearer <token>
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Неверный формат токена"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims := &utils.JWTClaims{}

		// Парсим токен с проверкой подписи
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Недействительный или просроченный токен"})
			c.Abort()
			return
		}

		// Успешная аутентификация — передаём userID в контекст
		c.Set("userID", uint(claims.UserID))
		c.Next()
	}
}
