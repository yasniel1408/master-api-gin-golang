package albums_ports

import "github.com/gin-gonic/gin"

type AlbumControllerInterface interface {
	GetAlbumsController(c *gin.Context) any
	NewAlbumsController(c *gin.Context) any
	GetAlbumController(c *gin.Context) any
	DeleteAlbumsController(c *gin.Context) any
}
