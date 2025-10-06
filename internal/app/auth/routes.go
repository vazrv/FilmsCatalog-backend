// internal/app/auth/routes.go
package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg gin.IRoutes, handler *AuthHandler) {
	rg.POST("/register", handler.Register)
	rg.POST("/login", handler.Login)
}
