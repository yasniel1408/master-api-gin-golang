package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	albums_infrastructure "github.com/yasniel1408/master-api-gin-golang/albums/infrastructure"
)

func main() {
	routers := gin.Default()
	// config cors
	routers.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	var albumsRouter = albums_infrastructure.AlbumController{}

	// rutas
	routers.GET("/albums", albumsRouter.GetAlbumsController)
	routers.POST("/albums", albumsRouter.NewAlbumsController)
	routers.GET("/albums/:id", albumsRouter.GetAlbumController)
	routers.DELETE("/albums/:id", albumsRouter.DeleteAlbumsController)
	routers.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// config server
	server := &http.Server{
		Addr:    ":8000",
		Handler: routers,
	}

	// run server
	server.ListenAndServe()
}
