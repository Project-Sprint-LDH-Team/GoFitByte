package services

import (
	"fit-byte-go/internal/models"
	"fit-byte-go/internal/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// update user data
func (s *UserService) UpdateUser(userID string, user *models.User) error {
	return s.repo.UpdateUser(userID, user)
}

// get user data
func (s *UserService) GetUserByID(userID string) (*models.User, error) {
	return s.repo.GetUserByID(userID)
}
