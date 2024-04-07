package auth_application

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strings"
	"time"
)

func generateToken(userID uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expira en 24 horas

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func generateRefreshToken(userID uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix() // Token de actualización expira en 30 días

	refreshTokenString, err := token.SignedString([]byte(os.Getenv("JWT_REFRESH_KEY")))
	if err != nil {
		return "", err
	}

	return refreshTokenString, nil
}

func getUserIdByToken(refreshTokenString string) (uint, error) {
	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Método de firma inesperado: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["userID"].(float64))
		return userID, nil
	} else {
		return 0, fmt.Errorf("Token de actualización inválido")
	}
}

func VerifyToken(token string) bool {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Método de firma inesperado: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return false
	}
	return parsedToken.Valid
}

func GetTokenWihoutBearer(tokenAndBearer string) (string, error) {
	splitToken := strings.Split(tokenAndBearer, " ")

	// verificar que exista el Bearer delante del token
	if splitToken[0] != "Bearer" {
		return "", errors.New("Authorization header format must be Bearer {token}")
	}
	token := strings.TrimSpace(splitToken[1])
	return token, nil
}
