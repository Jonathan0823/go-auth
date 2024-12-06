package user

import "go-auth/internal/models"

type UserService interface {
	Register(user *models.User) error
	Login(email, password string) (string, error)
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

func (s *userservice) Register(user *models.User) error {
	return s.repo.CreateUser(user)
}

func (s *userservice) GetUserByID(id int) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userservice) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.GetUserByEmail(email)
}