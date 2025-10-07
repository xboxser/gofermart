package service

import (
	"context"
	"gophermart/internal/user/model"
	"gophermart/internal/user/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Login(apiUser model.APIUser) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	userRepository := repository.NewUserRepository(ctx)
	user, err := userRepository.GetUserForLogin(apiUser.Login)
	if user.ID == 0 || err != nil {
		return -1, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(apiUser.Password))
	if err != nil {
		return user.ID, err
	}
	return user.ID, nil
}
