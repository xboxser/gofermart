package service

import (
	"gophermart/internal/balance/repository"
	"gophermart/internal/config/db"
	"gophermart/internal/withdraw/model"
)

type WithdrawalsService struct {
}

func NewWithdrawalsService() *WithdrawalsService {
	return &WithdrawalsService{}
}

func (s *WithdrawalsService) Withdrawals(userID int) ([]model.ApiWithdraw, error) {
	repository := repository.NewWithdrawRepository(db.GetDB())

	withdrawals, err := repository.GetWithdraws(userID)
	if err != nil {
		return nil, err
	}

	return withdrawals, nil
}
