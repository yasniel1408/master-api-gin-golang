package auth_middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	auth_application "github.com/yasniel1408/master-api-gin-golang/auth/application"
	shared_errors "github.com/yasniel1408/master-api-gin-golang/shared/models"
	"net/http"
)

func VerifyTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenAndBearer := c.GetHeader("Authorization")

		token, err := auth_application.GetTokenWihoutBearer(tokenAndBearer)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, shared_errors.ErrorStruct{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			})
			return
		}

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, shared_errors.ErrorStruct{
				Message: "Token no proporcionado",
				Status:  http.StatusBadRequest,
			})
			return
		}

		// verificar el token
		var isValid = auth_application.VerifyToken(token)
		fmt.Println("isValid", isValid)

		if !isValid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, shared_errors.ErrorStruct{
				Message: "Token invalid",
				Status:  http.StatusBadRequest,
			})
			return
		} else {
			// Continuar con el manejo de la solicitud
			c.Next()
		}
	}
}
