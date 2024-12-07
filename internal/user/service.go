package user

import "go-auth/internal/models"

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type userservice struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *userservice {
	return &userservice{
		repo: repo,
	}
}

func (s *userservice) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *userservice) GetUserByID(id int) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userservice) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.GetUserByEmail(email)
}