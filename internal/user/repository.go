package user

import (
	"database/sql"
	"go-auth/internal/models"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

type userrepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userrepository {
	return &userrepository{
		db: db,
	}
}

func (r *userrepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	query := `SELECT * FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
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



func (r *userrepository) UpdateUser(user *models.User) error {
	query := `UPDATE users SET username = $1, email = $2, WHERE id = $3`

	_, err := r.db.Exec(query, user.Username, user.Email, user.ID)
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
