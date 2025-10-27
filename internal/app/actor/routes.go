// internal/app/actor/routes.go
package actor

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg gin.IRoutes, handler *Handler) {
	rg.GET("/top", handler.GetTop)
}