package repositories

import (
	"go-learn/repositories/auth_repo"
	"go-learn/repositories/product_repo"
)

type Repo struct {
	AuthRepo    auth_repo.AuthContract
	ProductRepo product_repo.ProductContract
}

func NewRepo() *Repo {
	return &Repo{
		AuthRepo:    auth_repo.NewAuthRepositories(),
		ProductRepo: product_repo.NewProductRepositories(),
	}
}
