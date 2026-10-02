package repository

import (
	"context"

	"veerai/internal/database"
	"veerai/internal/models"
)


func GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}

	err := database.DB.QueryRow(
		context.Background(),
		`SELECT id, username, password
		 FROM users
		 WHERE username = $1`,
		username,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func CreateUser(user models.User) error {

	_, err := database.DB.Exec(
		context.Background(),
		`INSERT INTO users(username,password)
		 VALUES($1,$2)`,
		user.Username,
		user.Password,
	)

	return err
}
