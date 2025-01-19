package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssitko/hex-domain/internal/handlers"
)

func RegisterAlbumHandlers(router *gin.Engine, handler *handlers.AlbumHandler) *gin.RouterGroup {
	albumRouter := router.Group("/v1")
	{
		// Album routes
		albumRouter.GET("/albums", handler.GetAlbums)
		albumRouter.GET("/albums/:id", handler.GetAlbumByID)
		albumRouter.POST("/albums", handler.CreateAlbum)
		albumRouter.PUT("/albums", handler.UpdateAlbum)
		albumRouter.DELETE("/albums/:id", handler.DeleteAlbum)
	}
	return albumRouter
}
