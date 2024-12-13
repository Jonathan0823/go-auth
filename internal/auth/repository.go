package auth

import (
	"database/sql"
	"go-auth/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type AuthRepository interface {
	Register(user models.User) error
	Validate(email, password string) (bool, error)
	IsUserExists(email string) (bool)
}

type authrepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *authrepository {
	return &authrepository{
		db: db,
	}
}

func (r *authrepository) Validate(email, password string) (bool, error) {
	query := `SELECT password FROM users WHERE email = $1`

	var hashedPassword string
	err := r.db.QueryRow(query, email).Scan(&hashedPassword)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, nil
	}
	return true, nil
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


func (r *authrepository) IsUserExists(email string) (bool) {
	query := `SELECT email FROM users WHERE email = $1`

	var userEmail string
	err := r.db.QueryRow(query, email).Scan(&userEmail)
	if err != nil {
		return false
	}
	return true
}