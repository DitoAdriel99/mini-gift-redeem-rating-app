package auth_service

import (
	"fmt"
	"go-learn/entities"
	"go-learn/library/hashing"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func (s *_Service) Login(payload *entities.Login) (*entities.LoginResponse, error) {

	respUser, err := s.repo.AuthRepo.Checklogin(payload)
	if err != nil {
		log.Println("check login is error : ", err)
		return nil, err
	}

	if ok := hashing.CheckPasswordHash(payload.Password, respUser.Password); !ok {
		log.Println("check password hashing is : ", ok)

		return nil, fmt.Errorf("password is not match!")
	}

	expFromEnv := os.Getenv("EXPIRED")
	expire, _ := strconv.Atoi(expFromEnv)
	expirationTime := time.Now().Add(time.Duration(expire) * time.Hour)

	claims := &entities.Claims{
		Username: respUser.FullName,
		Email:    respUser.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(entities.JWTKEY)
	if err != nil {
		log.Println("Signed String error : ", err)
		return nil, err
	}

	return &entities.LoginResponse{FullName: respUser.FullName, Email: respUser.Email, Token: tokenString}, nil
}
