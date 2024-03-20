package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Year: 1957},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Year: 1962},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Year: 1954},
	{ID: "4", Title: "Time Out", Artist: "Dave Brubeck", Year: 1959},
	{ID: "5", Title: "Somethin' Else", Artist: "Cannonball Adderley", Year: 1958},
	{ID: "6", Title: "The Sidewinder", Artist: "Lee Morgan", Year: 1963},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

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

	// rutas
	routers.GET("/albums", getAlbums)
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
