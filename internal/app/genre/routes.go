// internal/app/genre/routes.go
package genre

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg gin.IRoutes, handler *Handler) {
	rg.GET("", handler.GetAll)
}