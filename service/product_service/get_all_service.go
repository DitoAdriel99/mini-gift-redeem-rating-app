package product_service

import (
	"go-learn/entities"
	"go-learn/library/meta"
	"log"
)

func (s *_Service) GetAll(m *meta.Metadata) ([]entities.Product, error) {
	resp, err := s.repo.ProductRepo.GetAll(m)
	if err != nil {
		log.Println("Get All is error : ", err)
		return nil, err
	}

	return resp, nil
}
