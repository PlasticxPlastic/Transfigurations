package database

import (
	"database/sql"
	"errors"
	"github.com/transfiguration/internal/models"
)

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := DB.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).
		Scan(&user.ID, &user.Username, &user.Password)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	
	return &user, nil
} 