package service

import (
	"context"
	"gophermart/internal/balance/repository"
	"gophermart/internal/config/db"
	"gophermart/internal/withdraw/model"
	"time"
)

type WithdrawalsService struct {
}

func NewWithdrawalsService() *WithdrawalsService {
	return &WithdrawalsService{}
}

func (s *WithdrawalsService) Withdrawals(userID int) ([]model.APIWithdraw, error) {
	repository := repository.NewWithdrawRepository(db.GetDB())
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	withdrawals, err := repository.GetWithdraws(ctx, userID)
	if err != nil {
		return nil, err
	}

	return withdrawals, nil
}
