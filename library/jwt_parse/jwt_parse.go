package jwt_parse

import (
	"fmt"
	"go-learn/entities"
	"log"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func GetEmailFromToken(param string) (string, error) {
	tokenString := strings.TrimPrefix(param, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, &entities.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return entities.JWTKEY, nil
	})

	if err != nil {
		return "", fmt.Errorf("Error parsing token: %s", err)
	}

	if claims, ok := token.Claims.(*entities.Claims); ok && token.Valid {
		log.Println("Token is valid")
		return claims.Email, nil
	} else {
		log.Println("Token is Invalid")
		return "", fmt.Errorf("Invalid token")
	}
}
