package auth_infrastructure_in

import (
	"fmt"
	"github.com/gin-gonic/gin"
	auth_application "github.com/yasniel1408/master-api-gin-golang/auth/application"
	auth_ports "github.com/yasniel1408/master-api-gin-golang/auth/domain/ports"
	auth_infrastructure_dto "github.com/yasniel1408/master-api-gin-golang/auth/infrastructure/in-adapters/dto"
	shared_errors "github.com/yasniel1408/master-api-gin-golang/shared/models"
	"net/http"
)

type AuthController struct {
	auth_ports.AuthControllerInterfacePort
	auth_application.AuthService
}

func (e *AuthController) RegisterController(c *gin.Context) {
	var newRegister auth_infrastructure_dto.AuthDTO
	c.BindJSON(&newRegister)

	var err = newRegister.Validate()
	if err != nil {
		fmt.Println(err)
		var errorStruct = shared_errors.ErrorStruct{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
		c.IndentedJSON(http.StatusBadRequest, errorStruct)
		return
	}

	c.IndentedJSON(http.StatusOK, e.RegisterService(newRegister))
}

func (e *AuthController) LoginController(c *gin.Context) {
	var loginData auth_infrastructure_dto.AuthDTO
	c.BindJSON(&loginData)

	var err = loginData.Validate()
	if err != nil {
		fmt.Println(err)
		var errorStruct = shared_errors.ErrorStruct{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
		c.IndentedJSON(http.StatusBadRequest, errorStruct)
		return
	}

	c.IndentedJSON(http.StatusOK, e.LoginService(loginData))
}

func (e *AuthController) WhoiamController(c *gin.Context) {
	tokenAndBearer := c.GetHeader("Authorization")
	token, err := auth_application.GetTokenWihoutBearer(tokenAndBearer)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, shared_errors.ErrorStruct{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
	}

	c.IndentedJSON(http.StatusOK, e.WhoiamService(token))
}
