package auth_service

import (
	"fmt"
	"go-learn/entities"
	"log"
)

func (s *_Service) UpdateStatus(payload *entities.StatusPayload) error {
	respUser, err := s.repo.AuthRepo.ValidateUserId(payload.ID)
	if err != nil {
		log.Println("Validate User Id is error : ", err)
		return err
	}

	if respUser.IsActive == payload.IsActive {
		return fmt.Errorf(fmt.Sprintf("User Is Alread %t", payload.IsActive))
	}

	if err = s.repo.AuthRepo.UpdateStatusUser(payload.ID, payload.IsActive); err != nil {
		return err
	}

	return nil
}
