package albums_infrastructure_in

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	albums_infrastructure_dto "github.com/yasniel1408/master-api-gin-golang/albums/infrastructure/in-adapters/dto"
)

func TestGetAlbumsController(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Create a new instance of the AlbumController
	albumController := &AlbumController{}

	// Define a route for the GetAlbumsController function
	router.GET("/albums", albumController.GetAlbumsController)

	// Create a new HTTP request to the /albums route
	req, err := http.NewRequest("GET", "/albums", nil)
	assert.NoError(t, err)

	// Create a new HTTP response recorder
	recorder := httptest.NewRecorder()

	// Perform the HTTP request
	router.ServeHTTP(recorder, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestNewAlbumsController(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Create a new instance of the AlbumController
	albumController := &AlbumController{}

	// Define a route for the NewAlbumsController function
	router.POST("/albums", albumController.NewAlbumsController)

	// Create a new HTTP request to the /albums route
	album := albums_infrastructure_dto.AlbumDTO{
		Title:  "Test Album",
		Artist: "Test Artist",
		Year:   2021,
	}
	albumJSON, _ := json.Marshal(album)
	req, err := http.NewRequest("POST", "/albums", bytes.NewBuffer(albumJSON))
	assert.NoError(t, err)

	// Create a new HTTP response recorder
	recorder := httptest.NewRecorder()

	// Perform the HTTP request
	router.ServeHTTP(recorder, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)
}
