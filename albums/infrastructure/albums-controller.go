package albums_infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	albums_application "github.com/yasniel1408/master-api-gin-golang/albums/application"
	albums_models "github.com/yasniel1408/master-api-gin-golang/albums/domine/models"
	albums_ports "github.com/yasniel1408/master-api-gin-golang/albums/domine/ports"
)

type AlbumController struct {
	albums_ports.AlbumControllerInterface
}

func (e *AlbumController) GetAlbumsController(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums_application.GetAlbumsService())
}

func (e *AlbumController) NewAlbumsController(c *gin.Context) {
	var newAlbum albums_models.Album

	c.BindJSON(&newAlbum)

	var albums = albums_application.SaveAlbumsService(newAlbum)

	c.IndentedJSON(http.StatusOK, albums)
}

func (e *AlbumController) GetAlbumController(c *gin.Context) {
	albumID := c.Param("id")

	var album, error = albums_application.GetAlbumService(albumID)

	if error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func (e *AlbumController) DeleteAlbumsController(c *gin.Context) {
	var albumID string
	c.BindJSON(&albumID)

	var albums, error = albums_application.DeleteAlbumsService(albumID)

	if error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}
