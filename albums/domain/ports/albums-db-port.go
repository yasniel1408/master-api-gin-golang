package albums_ports

import (
	"github.com/gin-gonic/gin"
	albums_models "github.com/yasniel1408/master-api-gin-golang/albums/domain/models"
)

type AlbumDBInterfacePort interface {
	ListAlbums(c *gin.Context) any
	SaveAlbum(c *gin.Context) any
	GetAlbum(c *gin.Context) (albums_models.Album, error)
	DeleteAlbum(c *gin.Context) ([]albums_models.Album, error)
}
