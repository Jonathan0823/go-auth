package auth

import (
	"database/sql"
	"go-auth/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type AuthRepository interface {
	Login(email, password string) (string, error)
	Register(user models.User) error
}

type authrepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *authrepository {
	return &authrepository{
		db: db,
	}
}

func (r *authrepository) Login(email, password string) (string, error) {
	panic("not implemented") // TODO: Implement
}

func (r *authrepository) Register(user models.User) error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, user.Username, user.Email, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}
