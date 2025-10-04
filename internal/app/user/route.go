// internal/app/user/routes.go
package user

import "github.com/gin-gonic/gin"

// RegisterRoutes регистрирует маршруты для пользователей
func RegisterRoutes(rg gin.IRoutes, handler *UserHandler) {
	rg.GET("/profile", handler.GetProfile)
}
