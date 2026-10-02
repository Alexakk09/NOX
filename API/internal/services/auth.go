package services

import (
	"errors"

	"veerai/internal/models"
	"veerai/internal/repository"
	"veerai/internal/utils"
)

func Login(req models.LoginRequest) (*models.LoginResponse, error) {

	user, err := repository.GetUserByUsername(req.Username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	err = utils.CheckPassword(req.Password, user.Password)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	token, err := utils.GenerateJWT(user.ID, user.Username)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Token: token,
	}, nil
}

func Register(req models.LoginRequest) error {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := models.User{
		Username: req.Username,
		Password: hashedPassword,
	}

	return repository.CreateUser(user)
}
