package auth_infrastructure_db

import (
	auth_models "github.com/yasniel1408/master-api-gin-golang/auth/domain/models"
	auth_ports "github.com/yasniel1408/master-api-gin-golang/auth/domain/ports"
	"github.com/yasniel1408/master-api-gin-golang/shared"
)

type AuthDB struct {
	auth_ports.AlbumDBInterfacePort
}

func (e *AuthDB) CreateUser(user auth_models.User) auth_models.User {
	shared.ConnectGorm().Create(&user)
	return user
}

func (e *AuthDB) GetUserByEmail(email string) auth_models.User {
	var user auth_models.User
	shared.ConnectGorm().Where("email = ?", email).First(&user)
	return user
}

func (e *AuthDB) GetUserById(id int) auth_models.User {
	var user auth_models.User
	shared.ConnectGorm().Where("id = ?", id).First(&user)
	return user
}
