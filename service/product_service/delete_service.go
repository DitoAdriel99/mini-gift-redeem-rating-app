package product_service

import (
	"log"

	"github.com/google/uuid"
)

func (s *_Service) Delete(id uuid.UUID) error {
	_, err := s.repo.ProductRepo.Detail(id)
	if err != nil {
		log.Println("Detail is error : ", err)
		return err
	}

	if err := s.repo.ProductRepo.Delete(id); err != nil {
		log.Println("Delete is error : ", err)
		return err
	}
	return nil
}
