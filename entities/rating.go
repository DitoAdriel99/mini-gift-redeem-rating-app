package entities

import (
	"go-learn/library/errbank"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

const (
	ErrAlreadyRating errbank.Error = "Already Rating this Item!"
)

type RatingRequired struct {
	ID        uuid.UUID  `json:"id"`
	ProductID uuid.UUID  `json:"product_id"`
	Email     string     `json:"email_user"`
	Rating    float64    `json:"rating"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
type PayloadRating struct {
	Rating float64 `json:"rating"`
}

func (l PayloadRating) Validate() error {
	return validation.ValidateStruct(
		&l,
		validation.Field(&l.Rating, validation.Required),
	)
}
