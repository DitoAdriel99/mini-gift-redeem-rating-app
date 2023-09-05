package product_repo

import (
	"database/sql"
	"go-learn/config"
	"go-learn/entities"
	"go-learn/library/meta"

	"github.com/google/uuid"
)

type _ProductRepoImp struct {
	conn *sql.DB
}

type ProductContract interface {
	Create(pr *entities.Product) error
	Detail(id uuid.UUID) (*entities.Product, error)
	GetAll(m *meta.Metadata) ([]entities.Product, error)
	Update(pr *entities.Product) error
	Delete(id uuid.UUID) error
	Redeem(rq *entities.RedeemRequired) error
	Rating(rq *entities.RatingRequired) error
}

func NewProductRepositories() ProductContract {
	conn, err := config.DBConn()
	if err != nil {
		panic(err)
	}

	return &_ProductRepoImp{
		conn: conn,
	}
}
