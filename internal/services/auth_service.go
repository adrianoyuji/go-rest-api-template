package services

import (
	"errors"

	"github.com/adrianoyuji/go-rest-api-template/internal/models"
	"github.com/adrianoyuji/go-rest-api-template/internal/repositories"
)

type AuthService struct {
	Users *repositories.UserRepository
}

func NewAuthService(u *repositories.UserRepository) *AuthService {
	return &AuthService{Users: u}
}

func (s *AuthService) Register(user *models.User) error {
	// NOTE: hash password in real app
	return s.Users.Create(user)
}

func (s *AuthService) Authenticate(email, password string) (*models.User, error) {
	user, err := s.Users.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}
