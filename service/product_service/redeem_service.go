package product_service

import (
	"go-learn/entities"
	"go-learn/library/jwt_parse"
	"log"
	"time"

	"github.com/google/uuid"
)

func (s *_Service) Redeem(id uuid.UUID, payload *entities.PayloadRedeem, bearer string) error {
	var (
		time     = time.Now().Local()
		newId, _ = uuid.NewUUID()
	)
	emailUser, err := jwt_parse.GetEmailFromToken(bearer)
	if err != nil {
		log.Println("Get Email From Token error: ", err)
		return err
	}

	data, err := s.repo.ProductRepo.Detail(id)
	if err != nil {
		log.Println("Detail error: ", err)
		return err
	}

	redeemQty := data.Qty - payload.Qty
	if redeemQty < 0 {
		log.Println("Request Over Limit error: ", payload.Qty)
		return entities.ErrOverLimit
	}

	requestRedeem := entities.RedeemRequired{
		ID:        newId,
		ProductID: data.ID,
		QtyReq:    payload.Qty,
		QtyAfter:  redeemQty,
		Email:     emailUser,
		CreatedAt: time,
		UpdatedAt: &time,
	}

	if err := s.repo.ProductRepo.Redeem(&requestRedeem); err != nil {
		log.Println("Redeem error: ", err)
		return err
	}

	return nil
}
