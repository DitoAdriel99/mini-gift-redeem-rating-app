package service

import (
	"go-learn/repositories"
	"go-learn/service/auth_service"
	"go-learn/service/product_service"
)

type Service struct {
	AuthService    auth_service.Contract
	ProductService product_service.Contract
}

func NewService(repo *repositories.Repo) *Service {
	return &Service{
		AuthService:    auth_service.NewAuthService(repo),
		ProductService: product_service.NewProductService(repo),
	}
}
