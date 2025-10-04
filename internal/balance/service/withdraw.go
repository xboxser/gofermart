package service

import (
	"fmt"
	"gophermart/internal/balance/model"
	"gophermart/internal/balance/repository"
	"net/http"
)

type Withdraw struct {
	repositoryBalance  *repository.BalanceRepository
	repositoryWithdraw *repository.WithdrawRepository
}

func NewWithdraw() *Withdraw {
	return &Withdraw{
		repositoryBalance: repository.NewBalanceRepository(),
	}
}

func (w *Withdraw) Withdraw(userID int, model model.Withdraw) (int, error) {
	balance, err := w.repositoryBalance.GetBalanceUser(userID)
	if err != nil {
		return 0, fmt.Errorf("error get balance user")
	}

	if balance.Current < model.Sum {
		return http.StatusPaymentRequired, nil
	}
	fmt.Println("balance", balance)
	fmt.Println("model", model)
	return 1, nil
}
