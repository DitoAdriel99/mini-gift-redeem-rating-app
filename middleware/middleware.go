package middleware

import (
	"encoding/json"
	"fmt"
	"go-learn/entities"
	"go-learn/repositories"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type TokenValidator struct {
	Repo repositories.Repo
}

func NewTokenValidator(repo repositories.Repo) *TokenValidator {
	return &TokenValidator{
		Repo: repo,
	}
}

func (m *TokenValidator) ValidateTokenMiddleware(allowedRoles ...string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the token from the request headers or cookies
			tokenString := getTokenFromRequest(r)

			// Validate the token and get the user's role
			role, err := validateToken(tokenString, m.Repo)
			if err != nil || !isValidRole(role, allowedRoles) {
				errorResponse := struct {
					Message string `json:"message"`
				}{
					Message: "This User Is Unauthorized",
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(errorResponse)
				return
			}

			// Call the next handler if the role is valid
			next.ServeHTTP(w, r)
		})
	}
}

func isValidRole(role string, allowedRoles []string) bool {
	for _, allowedRole := range allowedRoles {
		if role == allowedRole {
			return true
		}
	}
	return false
}

func getTokenFromRequest(r *http.Request) string {
	return r.Header.Get("Authorization")
}

func validateToken(param string, repo repositories.Repo) (string, error) {
	tokenString := strings.TrimPrefix(param, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, &entities.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return entities.JWTKEY, nil
	})

	if err != nil {
		return "", fmt.Errorf("Error parsing token: %s", err)
	}

	if claims, ok := token.Claims.(*entities.Claims); ok && token.Valid {
		userData, err := repo.AuthRepo.ValidateUser(claims.Email)
		if err != nil {
			return "", err
		}
		return userData.Role, nil
	} else {
		log.Println("Token is Invalid")
		return "", fmt.Errorf("Invalid token")
	}
}
