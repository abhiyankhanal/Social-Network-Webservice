package utilities

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(username string) string {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET")))
	return tokenString
}
