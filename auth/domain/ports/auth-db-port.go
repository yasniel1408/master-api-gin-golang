package auth_ports

import (
	"github.com/gin-gonic/gin"
	auth_models "github.com/yasniel1408/master-api-gin-golang/auth/domain/models"
)

type AlbumDBInterfacePort interface {
	CreateUser(c *gin.Context) auth_models.User
	GetUserByEmail(c *gin.Context) auth_models.User
	GetUserById(c *gin.Context) auth_models.User
}
