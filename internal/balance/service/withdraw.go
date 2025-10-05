package service

import (
	"context"
	"fmt"
	"gophermart/internal/balance/model"
	"gophermart/internal/balance/repository"
	"gophermart/internal/config/db"
	"log"
	"net/http"
	"time"
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

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	db := db.GetDB()
	tx, err := db.Begin(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback(ctx)

	err = w.repositoryBalance.WithdrawBalanceUser(tx, ctx, userID, model.Sum)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error WithdrawBalanceUser")
	}

	err = w.repositoryWithdraw.AddWithdraw(tx, ctx, userID, model.Sum, model.OrderID)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, fmt.Errorf("error AddWithdraw")
	}

	tx.Commit(ctx)
	return http.StatusOK, nil
}
