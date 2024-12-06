package user

import (
	"database/sql"
	"go-auth/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetUserByID(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	UpdatePassword(id int, password string) error
	DeleteUser(id int) error
	VerifyUser(email, password string) (*models.User, error)
}

type userrepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userrepository {
	return &userrepository{
		db: db,
	}
}

func (r *userrepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userrepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userrepository) CreateUser(user *models.User) error {
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

func (r *userrepository) UpdateUser(user *models.User) error {
	query := `UPDATE users SET username = $1, email = $2, WHERE id = $3`

	_, err := r.db.Exec(query, user.Username, user.Email, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *userrepository) UpdatePassword(id int, password string) error {
	query := `UPDATE users SET password = $1 WHERE id = $2`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, hashedPassword, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *userrepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *userrepository) VerifyUser(email, password string) (*models.User, error) {
	user, err := r.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}