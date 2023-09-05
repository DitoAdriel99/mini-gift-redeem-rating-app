package auth_service

import (
	"go-learn/entities"
	"go-learn/repositories"
)

type Contract interface {
	Login(payload *entities.Login) (*entities.LoginResponse, error)
	Register(payload *entities.RegisterPayload) (*entities.Register, error)
	UpdateStatus(payload *entities.StatusPayload) error
}

type _Service struct {
	repo *repositories.Repo
}

func NewAuthService(repo *repositories.Repo) Contract {
	return &_Service{repo}
}
