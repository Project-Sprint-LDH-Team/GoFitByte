package services

import (
	"errors"
	"fit-byte-go/internal/models"
	"fit-byte-go/internal/repositories"
	"fit-byte-go/internal/utils"
)

type AuthService struct {
	repo *repositories.AuthRepository
}

func NewAuthService(repo *repositories.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

// register new user
func (s *AuthService) Register(user *models.AuthRequest) error {
	// check email already registered
	existingUser, err := s.repo.FindUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("email already exists")
	}
	// Hash password before to store on db
	user.Password = utils.HashPassword(user.Password)
	// Store user on db
	return s.repo.Register(user)
}

// Check credentials and return token
func (s *AuthService) Login(email, password string) (*models.AuthResponse, error) {
	// Find user based on email
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("email not found")
	}
	// Password verification
	if !utils.VerifyPassword(user.Password, password) {
		return nil, errors.New("invalid password")
	}

	// Generate token
	token := utils.GenerateToken()
	return &models.AuthResponse{
		Email: user.Email,
		Token: token,
	}, nil
}
