package service

import (
	"gophermart/internal/user/model"

	"golang.org/x/crypto/bcrypt"
)

func Login(apiUser model.APIUser) (int, error) {

	user, err := GetUserForLogin(apiUser.Login)
	if user.ID == 0 || err != nil {
		return -1, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(apiUser.Password))
	if err != nil {
		return user.ID, err
	}
	return user.ID, nil
}
