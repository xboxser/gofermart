package service

import (
	"context"
	"fmt"
	balanceRepository "gophermart/internal/balance/repository"
	"gophermart/internal/config/db"
	"gophermart/internal/user/model"
	"gophermart/internal/user/repository"
	"log"
	"time"
)

func Register(apiUser model.APIUser) (int, error) {
	db := db.GetDB()

	user, err := GetUserForLogin(apiUser.Login)

	if user.ID != 0 || err != nil {
		return -1, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	userRepository := repository.NewUserRepository(ctx)

	tx, err := db.Begin(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback(ctx) // Если не отработает commit, то откатываем в любом случае

	userID, err := userRepository.RegisterUser(tx, apiUser.Login, apiUser.Password)
	if err != nil {
		return -1, fmt.Errorf("error registering user: %v", err)
	}

	if userID == -1 {
		return -1, fmt.Errorf("error registering user")
	}

	balanceRepository := balanceRepository.NewBalanceRepository()
	err = balanceRepository.CreateBalanceUser(tx, ctx, userID)
	if err != nil {
		return -1, fmt.Errorf("error creating balance: %v", err)
	}

	tx.Commit(ctx)

	return userID, nil
}
