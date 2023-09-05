package product_service

import (
	"go-learn/entities"
	"log"
	"time"

	"github.com/google/uuid"
)

func (s *_Service) Create(payload *entities.Product) error {
	var (
		time  = time.Now().Local()
		id, _ = uuid.NewUUID()
	)
	payload.ID = id
	payload.CreatedAt = time
	payload.UpdatedAt = &time

	if err := s.repo.ProductRepo.Create(payload); err != nil {
		log.Println("create product is error : ", err)
		return err
	}

	return nil
}
