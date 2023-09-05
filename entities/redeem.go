package entities

import (
	"time"

	"go-learn/library/errbank"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

const (
	ErrOverLimit   errbank.Error = "Request Over Limit!"
	ErrNeverRedeem errbank.Error = "You Never Redeem This Item!"
)

type RedeemRequired struct {
	ID        uuid.UUID  `json:"id"`
	ProductID uuid.UUID  `json:"product_id"`
	QtyReq    int        `json:"quantity_request"`
	QtyBefore int        `json:"quantity_before"`
	QtyAfter  int        `json:"quantity_after"`
	Email     string     `json:"email_user"`
	Rating    float64    `json:"rating"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type PayloadRedeem struct {
	Qty int `json:"quantity"`
}

func (l PayloadRedeem) Validate() error {
	return validation.ValidateStruct(
		&l,
		validation.Field(&l.Qty, validation.Required),
	)
}
