package main

import (
	"fmt"
	"github.com/yasniel1408/master-api-gin-golang/shared"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	albums_infrastructure_in "github.com/yasniel1408/master-api-gin-golang/albums/infrastructure/in-adapters"
	auth_infrastructure_in "github.com/yasniel1408/master-api-gin-golang/auth/infrastructure/in-adapters"
)

func init() {
	// load envs
	shared.LoadEnvs()
	// Conectar a DB
	shared.ConnectGorm()
}

func main() {
	routers := gin.Default()
	var ginmode = os.Getenv("GIN_MODE")
	gin.SetMode(ginmode)
	// Config Middleware
	// CORS
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
	// Logger
	routers.Use(gin.Logger())

	// rutas
	albums_infrastructure_in.AlbumRouter(routers)
	auth_infrastructure_in.AuthRouter(routers)
	routers.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Config Server
	var port = ":" + os.Getenv("PORT")
	server := &http.Server{
		Addr:    port,
		Handler: routers,
	}

	// Run Server
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
