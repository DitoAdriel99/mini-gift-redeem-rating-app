package entities

import (
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
type Register struct {
	ID        uuid.UUID  `json:"id"`
	FullName  string     `json:"fullname"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type RegisterPayload struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l RegisterPayload) Validate() error {
	return validation.ValidateStruct(
		&l,
		validation.Field(&l.FullName, validation.Required),
		validation.Field(&l.Email, validation.Required),
		validation.Field(&l.Password, validation.Required),
	)
}

type StatusPayload struct {
	ID       uuid.UUID `json:"id"`
	IsActive bool      `json:"is_active"`
}

func (l StatusPayload) Validate() error {
	if l.IsActive != true && l.IsActive != false {
		return errors.New("IsActive must be true or false")
	}
	return validation.ValidateStruct(
		&l,
		validation.Field(&l.ID, validation.Required),
	)
}
