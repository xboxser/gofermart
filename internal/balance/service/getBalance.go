package service

import (
	"gophermart/internal/balance/model"
	"gophermart/internal/balance/repository"
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
	balance, err := g.repository.GetBalanceUser(g.userID)
	return balance, err
}
