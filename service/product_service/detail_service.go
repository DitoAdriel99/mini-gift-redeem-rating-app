package product_service

import (
	"go-learn/entities"
	"log"

	"github.com/google/uuid"
)

func (s *_Service) Detail(id uuid.UUID) (*entities.Product, error) {
	data, err := s.repo.ProductRepo.Detail(id)
	if err != nil {
		log.Println("Detail product is error : ", err)
		return nil, err
	}

	return data, nil
}
