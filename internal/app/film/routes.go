// internal/app/film/routes.go
package film

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg gin.IRoutes, handler *Handler) {
	rg.GET("/popular", handler.GetPopular)
}

// маршрут запускает хендлер
