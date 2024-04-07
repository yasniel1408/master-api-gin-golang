package auth_infrastructure_in

import (
	"github.com/gin-gonic/gin"
	auth_middlewares "github.com/yasniel1408/master-api-gin-golang/auth/infrastructure/in-adapters/middlewares"
)

func AuthRouter(routers *gin.Engine) {
	var authRouter = AuthController{}

	routers.POST("/login", authRouter.LoginController)
	routers.POST("/register", authRouter.RegisterController)
	routers.POST("/whoiam", auth_middlewares.VerifyTokenMiddleware(), authRouter.WhoiamController)
}
