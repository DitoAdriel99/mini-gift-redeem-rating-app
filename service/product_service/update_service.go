package product_service

import (
	"go-learn/entities"
	"log"
	"time"

	"github.com/google/uuid"
)

func (s *_Service) Update(id uuid.UUID, payload *entities.Product) error {
	var (
		time = time.Now().Local()
	)
	data, err := s.repo.ProductRepo.Detail(id)
	if err != nil {
		log.Println("get detail error : ", err)
		return err
	}

	payload.ID = data.ID

	if payload.Title == "" {
		payload.Title = data.Title
	}

	if payload.Description == "" {
		payload.Description = data.Description
	}
	if payload.Points == 0 {
		payload.Points = data.Points
	}
	if payload.Qty == 0 {
		payload.Qty = data.Qty
	}
	if payload.Image == "" {
		payload.Image = data.Image
	}
	if payload.Type == "" {
		payload.Type = data.Type
	}
	if payload.Banner == "" {
		payload.Banner = data.Banner
	}
	if payload.Info == "" {
		payload.Info = data.Info
	}

	payload.UpdatedAt = &time

	if err := s.repo.ProductRepo.Update(payload); err != nil {
		log.Println("Update error : ", err)
		return err
	}

	return nil
}
