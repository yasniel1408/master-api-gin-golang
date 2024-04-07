package albums_infrastructure_in

import (
	"net/http"

	"github.com/gin-gonic/gin"
	albums_application "github.com/yasniel1408/master-api-gin-golang/albums/application"
	albums_ports "github.com/yasniel1408/master-api-gin-golang/albums/domain/ports"
	albums_infrastructure_dto "github.com/yasniel1408/master-api-gin-golang/albums/infrastructure/in-adapters/dto"
)

type AlbumController struct {
	albums_ports.AlbumHttpInterfacePort
	albums_application.AlbumService
}

func (e *AlbumController) GetAlbumsController(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, e.GetAlbumsService())
}

func (e *AlbumController) NewAlbumsController(c *gin.Context) {
	var newAlbum albums_infrastructure_dto.AlbumDTO
	c.BindJSON(&newAlbum)

	var _, error = newAlbum.Validate()

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": error.Error()})
		return
	}

	var albums = e.SaveAlbumsService(newAlbum)

	c.IndentedJSON(http.StatusOK, albums)
}

func (e *AlbumController) GetAlbumController(c *gin.Context) {
	albumID := c.Param("id")

	var album, error = e.GetAlbumService(albumID)

	if error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func (e *AlbumController) DeleteAlbumsController(c *gin.Context) {
	albumID := c.Param("id")

	var albums, error = e.DeleteAlbumsService(albumID)

	if error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}
