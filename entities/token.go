package entities

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var JWTKEY = []byte(os.Getenv("KEY"))

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}
