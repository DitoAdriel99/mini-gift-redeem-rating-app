package product_service

import (
	"go-learn/entities"
	"go-learn/library/meta"
	"go-learn/repositories"

	"github.com/google/uuid"
)

type Contract interface {
	Create(payload *entities.Product) error
	Update(id uuid.UUID, payload *entities.Product) error
	Delete(id uuid.UUID) error
	Detail(id uuid.UUID) (*entities.Product, error)
	GetAll(m *meta.Metadata) ([]entities.Product, error)
	Redeem(id uuid.UUID, payload *entities.PayloadRedeem, bearer string) error
	Rating(id uuid.UUID, payload *entities.PayloadRating, bearer string) error
}

type _Service struct {
	repo *repositories.Repo
}

func NewProductService(repo *repositories.Repo) Contract {
	return &_Service{repo}
}
