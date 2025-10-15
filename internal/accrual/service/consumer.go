package service

import (
	"context"
	"errors"
	"gophermart/internal/accrual/model"
	repositoryBalance "gophermart/internal/balance/repository"
	"gophermart/internal/config/db"
	"gophermart/internal/logger"
	"gophermart/internal/order/repository"
	"log"
	"strconv"
	"time"
)

// Обрабатывает данные из потока
type Consumer struct {
	resultCh          chan model.OrderAccrual
	doneCh            chan struct{}
	repositoryOrder   *repository.OrderRepository
	repositoryBalance *repositoryBalance.BalanceRepository
	refundCh          chan int
}

func NewConsumer(resultCh chan model.OrderAccrual, doneCh chan struct{}, refundCh chan int) *Consumer {
	return &Consumer{
		repositoryOrder:   repository.NewOrderRepository(),
		repositoryBalance: repositoryBalance.NewBalanceRepository(),
		resultCh:          resultCh,
		refundCh:          refundCh,
	}
}

func (c *Consumer) Run() {
	sugar := logger.GetLogger()
	go func() {
		for order := range c.resultCh {
			select {
			case <-c.doneCh:
				return
			default:
			}
			numOrder, _ := strconv.Atoi(order.Order)
			switch order.Status {
			case model.OrderStatusProcessed:
				err := c.accrualPoints(order)
				sugar.Infof("set status processed: %v", err)
			case model.OrderStatusInvalid:
				c.repositoryOrder.SetStatusInvalid(order.Order)
				c.refundCh <- numOrder
			case model.OrderStatusRegistered:
			case model.OrderStatusProcessing:
				c.repositoryOrder.SetStatusProcessing(order.Order)

			}
		}
	}()
}

// начисляем баллы пользователю
func (c *Consumer) accrualPoints(OrderAccrual model.OrderAccrual) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// проверяем что данный заказ есть в системе и у него указан пользователь
	order, err := c.repositoryOrder.GetOrderByNumber(ctx, OrderAccrual.Order)
	if err != nil || order.ID == 0 {
		return errors.New("error get order")
	}

	if order.UserID == 0 {
		return errors.New("error get user")
	}

	db := db.GetDB()
	tx, err := db.Begin(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback(ctx) // Если не отработает commit, то откатываем в любом случае

	// Устанавливаем статус, что по заказу все расчеты прошли
	err = c.repositoryOrder.SetStatusProcessed(tx, ctx, OrderAccrual.Order, OrderAccrual.Accrual)
	if err != nil {
		return errors.New("error set status processed")
	}

	// Получаем заказ, для индетификации пользователя
	order, err = c.repositoryOrder.GetOrderByNumber(ctx, OrderAccrual.Order)
	if err != nil {
		return errors.New("error get order")
	}

	// Начисляем баллы на счет пользователя
	err = c.repositoryBalance.AddBalanceUser(tx, ctx, order.UserID, OrderAccrual.Accrual)
	if err != nil {
		return errors.New("error add balance")
	}

	tx.Commit(ctx)
	return nil
}
