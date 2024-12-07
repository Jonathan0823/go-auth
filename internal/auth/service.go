package auth

import "go-auth/internal/models"

type AuthService interface {
	Register(user models.User) error
}

type authservice struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) *authservice {
	return &authservice{
		repo: repo,
	}
}

func (s *authservice) Register(user models.User) error {
	return s.repo.Register(user)
}

