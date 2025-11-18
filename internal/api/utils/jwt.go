package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var SecretKey = []byte("ExampleSecretKey")

func GenerateToken(userName string) (string, error) {

	claims := jwt.MapClaims{
		"userName" : userName,
		"exp" : time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	tokenString, err := token.SignedString(SecretKey)
	return tokenString, err
}