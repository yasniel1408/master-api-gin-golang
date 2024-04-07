package auth_application

import (
	"fmt"
	auth_models "github.com/yasniel1408/master-api-gin-golang/auth/domain/models"
	auth_infrastructure_dto "github.com/yasniel1408/master-api-gin-golang/auth/infrastructure/in-adapters/dto"
	auth_infrastructure_db "github.com/yasniel1408/master-api-gin-golang/auth/infrastructure/out-adapters"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	auth_infrastructure_db.AuthDB
}

func (auth *AuthService) RegisterService(newRegister auth_infrastructure_dto.AuthDTO) auth_models.User {

	// encriptar el password
	var passEnc, err = bcrypt.GenerateFromPassword([]byte(newRegister.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	var user auth_models.User
	user.Email = newRegister.Email
	user.Password = string(passEnc)
	user.Name = ""
	user.Gender = ""

	var newUser = auth.CreateUser(user)

	fmt.Println(newUser)

	return newUser
}

func (auth *AuthService) LoginService(loginData auth_infrastructure_dto.AuthDTO) auth_models.Authentication {
	// recuperar el usuario por el email
	var user = auth.GetUserByEmail(loginData.Email)

	// comparar los passwords
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		fmt.Println("Contraseña incorrecta:", err)
		panic("contraseña inválida")
	}

	// crear el token y el refresh token
	var token, errT = generateToken(user.ID)
	if errT != nil {
		panic(errT)
	}

	var refreshToken, errRT = generateRefreshToken(user.ID)
	if errRT != nil {
		panic(errRT)
	}

	var response = auth_models.Authentication{
		Token:         token,
		Refresh_Token: refreshToken,
	}

	return response
}

func (auth *AuthService) WhoiamService(token string) auth_models.User {
	var id, err = getUserIdByToken(token)
	if err != nil {
		panic(err)
	}
	var user = auth.GetUserById(int(id))
	return user
}
