package auth_service

import (
	"go-learn/entities"
	"go-learn/library/hashing"
	"log"
	"time"

	"github.com/google/uuid"
)

func (s *_Service) Register(payload *entities.RegisterPayload) (*entities.Register, error) {
	var (
		newId, _ = uuid.NewUUID()
		time     = time.Now().Local()
	)
	if err := s.repo.AuthRepo.CheckEmail(payload.Email); err != nil {
		log.Println("Check Email error : ", err)
		return nil, err
	}

	hashedPass, err := hashing.HashPassword(payload.Password)
	if err != nil {
		log.Println("Hash Password error : ", err)
		return nil, err
	}

	dataRegister := entities.Register{
		ID:        newId,
		FullName:  payload.FullName,
		Email:     payload.Email,
		Password:  hashedPass,
		Role:      "user",
		CreatedAt: time,
		UpdatedAt: &time,
	}

	if err := s.repo.AuthRepo.Register(&dataRegister); err != nil {
		log.Println("Register error : ", err)
		return nil, err
	}

	return &dataRegister, nil
}
