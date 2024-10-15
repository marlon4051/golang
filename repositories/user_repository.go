package repositories

import (
	"database/sql"
	"task-api/models"
)

type UserRepository struct {
	DB *sql.DB
}

// get user by email
func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT id, username, email, password FROM users WHERE email = ?"
	row := r.DB.QueryRow(query, email)

	var user models.User
	err := row.Scan(&user.ID, &user.UserName, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No existe usuario con ese email
		}
		return nil, err
	}

	return &user, nil
}

// create user
func (r *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	_, err := r.DB.Exec(query, user.UserName, user.Email, user.Password)
	return err
}
