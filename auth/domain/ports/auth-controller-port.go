package auth_ports

import "github.com/gin-gonic/gin"

type AuthControllerInterfacePort interface {
	RegisterController(c *gin.Context) any
	LoginController(c *gin.Context) any
	WhoiamController(c *gin.Context) any
}
