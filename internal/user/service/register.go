package service

import (
	"context"
	"errors"
	"fmt"
	balanceRepository "gophermart/internal/balance/repository"
	"gophermart/internal/config/db"
	"gophermart/internal/user/model"
	"gophermart/internal/user/repository"
	"time"
)

func Register(apiUser model.APIUser) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	userRepository := repository.NewUserRepository()

	user, err := userRepository.GetUserForLogin(ctx, apiUser.Login)

	if err != nil {
		return 0, err
	}

	if user.ID != 0 {
		return 0, model.ErrLoginBusy
	}
	db := db.GetDB()
	tx, err := db.Begin(ctx)
	if err != nil {
		return 0, errors.New("error begin transaction")
	}
	defer tx.Rollback(ctx) // Если не отработает commit, то откатываем в любом случае

	userID, err := userRepository.RegisterUser(tx, ctx, apiUser.Login, apiUser.Password)
	if err != nil {
		return 0, fmt.Errorf("error registering user: %w", err)
	}

	balanceRepository := balanceRepository.NewBalanceRepository()
	err = balanceRepository.CreateBalanceUser(tx, ctx, userID)
	if err != nil {
		return 0, fmt.Errorf("error creating balance: %w", err)
	}

	tx.Commit(ctx)

	return userID, nil
}
