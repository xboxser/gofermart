package service

import (
	"context"
	"gophermart/internal/balance/model"
	"gophermart/internal/balance/repository"
	"time"
)

type GetBalance struct {
	userID     int
	repository *repository.BalanceRepository
}

func NewGetBalance(userID int) *GetBalance {
	return &GetBalance{
		userID:     userID,
		repository: repository.NewBalanceRepository(),
	}
}

func (g *GetBalance) GetBalanceUser() (model.Balance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	balance, err := g.repository.GetBalanceUser(ctx, g.userID)
	return balance, err
}
