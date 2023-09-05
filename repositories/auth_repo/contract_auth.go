package auth_repo

import (
	"database/sql"
	"go-learn/config"
	"go-learn/entities"

	"github.com/google/uuid"
)

type _AuthRepoImp struct {
	conn *sql.DB
}

type AuthContract interface {
	Checklogin(auth *entities.Login) (*entities.User, error)
	ValidateUser(email string) (*entities.User, error)
	CheckEmail(email string) error
	Register(rq *entities.Register) error
	UpdateStatusUser(id uuid.UUID, status bool) error
	ValidateUserId(id uuid.UUID) (*entities.User, error)
}

func NewAuthRepositories() AuthContract {
	conn, err := config.DBConn()
	if err != nil {
		panic(err)
	}

	return &_AuthRepoImp{
		conn: conn,
	}
}
