package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

var SecretKey []byte

func GenerateToken(userName string) (string, error) {

	if err := godotenv.Load("./.env"); err != nil {
		panic(err)
	}

	SecretKey = []byte(os.Getenv("secretkey"))

	claims := jwt.MapClaims{
		"userName": userName,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(SecretKey)
	return tokenString, err
}
