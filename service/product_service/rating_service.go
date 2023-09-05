package product_service

import (
	"go-learn/entities"
	"go-learn/library/jwt_parse"
	"log"
	"time"

	"github.com/google/uuid"
)

func (s *_Service) Rating(id uuid.UUID, payload *entities.PayloadRating, bearer string) error {
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

	ratingReq := entities.RatingRequired{
		ID:        newId,
		ProductID: data.ID,
		Email:     emailUser,
		Rating:    payload.Rating,
		CreatedAt: time,
		UpdatedAt: &time,
	}

	if err := s.repo.ProductRepo.Rating(&ratingReq); err != nil {
		log.Println("Detail error: ", err)
		return err
	}

	return nil
}
