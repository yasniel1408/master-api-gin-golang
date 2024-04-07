package albums_infrastructure_in

import (
	"github.com/gin-gonic/gin"
	auth_middlewares "github.com/yasniel1408/master-api-gin-golang/auth/infrastructure/in-adapters/middlewares"
)

func AlbumRouter(routers *gin.Engine) {
	var albumsRouter = AlbumController{}

	routers.GET("/albums", auth_middlewares.VerifyTokenMiddleware(), albumsRouter.GetAlbumsController)
	routers.POST("/albums", auth_middlewares.VerifyTokenMiddleware(), albumsRouter.NewAlbumsController)
	routers.GET("/albums/:id", auth_middlewares.VerifyTokenMiddleware(), albumsRouter.GetAlbumController)
	routers.DELETE("/albums/:id", auth_middlewares.VerifyTokenMiddleware(), albumsRouter.DeleteAlbumsController)
}
