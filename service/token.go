package service

import (
	"fmt"
)

func (s *AuthService) StoreSessionToken(adminID int, token string) error {
	err := s.Repo.AuthRepo.StoreSessionToken(adminID, token)
	if err != nil {
		return fmt.Errorf("failed to store session token: %v", err)
	}
	return nil
}
