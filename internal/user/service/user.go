package service

import (
	"context"
	"gophermart/internal/user/model"
	"gophermart/internal/user/repository"
	"time"
)

func GetUserForID(ID int) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	userRepository := repository.NewUserRepository()
	user, err := userRepository.GetUserForID(ctx, ID)

	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
