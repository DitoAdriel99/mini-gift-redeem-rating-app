package auth_repo

import (
	"database/sql"
	"fmt"
	"go-learn/entities"

	"github.com/google/uuid"
)

func (c *_AuthRepoImp) Checklogin(auth *entities.Login) (*entities.User, error) {
	query := `SELECT * FROM users WHERE email = $1`

	var object entities.User

	err := c.conn.QueryRow(query, auth.Email).Scan(
		&object.ID,
		&object.FullName,
		&object.Email,
		&object.Password,
		&object.Role,
		&object.CreatedAt,
		&object.UpdatedAt,
		&object.IsActive,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		err = fmt.Errorf("scanning activity objects: %w", err)
		return nil, err
	}

	return &object, nil

}

func (c *_AuthRepoImp) ValidateUser(email string) (*entities.User, error) {
	query := `SELECT * FROM users WHERE email = $1`

	var object entities.User

	err := c.conn.QueryRow(query, email).Scan(
		&object.ID,
		&object.FullName,
		&object.Email,
		&object.Password,
		&object.Role,
		&object.CreatedAt,
		&object.UpdatedAt,
		&object.IsActive,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		err = fmt.Errorf("scanning activity objects: %w", err)
		return nil, err
	}
	return &object, nil
}

func (c *_AuthRepoImp) CheckEmail(email string) error {
	query := `SELECT COUNT(*) FROM users WHERE email = $1`

	var count int
	err := c.conn.QueryRow(query, email).Scan(&count)
	if err != nil {
		err = fmt.Errorf("scanning activity objects: %w", err)
		return err
	}

	if count == 1 {
		err = fmt.Errorf("Email Already Used!")
		return err
	}

	return nil
}

func (c *_AuthRepoImp) Register(rq *entities.Register) error {
	queryInsert := `
		INSERT INTO users (id, fullname, email, password, role, created_at, updated_at, is_active)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	if _, err := c.conn.Exec(queryInsert, rq.ID, rq.FullName, rq.Email, rq.Password, rq.Role, rq.CreatedAt, rq.UpdatedAt, true); err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return err
	}

	return nil
}

func (c *_AuthRepoImp) UpdateStatusUser(id uuid.UUID, status bool) error {
	updateQuery := `
	UPDATE users
	SET
		is_active = $2
	WHERE
		id = $1
	`

	if _, err := c.conn.Exec(updateQuery, id, status); err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return err
	}
	return nil
}

func (c *_AuthRepoImp) ValidateUserId(id uuid.UUID) (*entities.User, error) {
	query := `SELECT * FROM users WHERE id = $1`

	var object entities.User

	err := c.conn.QueryRow(query, id).Scan(
		&object.ID,
		&object.FullName,
		&object.Email,
		&object.Password,
		&object.Role,
		&object.CreatedAt,
		&object.UpdatedAt,
		&object.IsActive,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		err = fmt.Errorf("scanning activity objects: %w", err)
		return nil, err
	}
	return &object, nil
}
