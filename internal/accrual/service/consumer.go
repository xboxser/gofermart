package service

import (
	"fmt"
	"gophermart/internal/accrual/model"
	"gophermart/internal/order/repository"
)

// Обрабатывает данные из потока
type Consumer struct {
	resultCh   chan model.OrderAccrual
	doneCh     chan struct{}
	Repository *repository.OrderRepository
}

func NewConsumer(resultCh chan model.OrderAccrual, doneCh chan struct{}) *Consumer {
	return &Consumer{
		Repository: repository.NewOrderRepository(),
		resultCh:   resultCh,
	}
}

func (c *Consumer) Run() {
	go func() {

		for order := range c.resultCh {
			select {
			case <-c.doneCh:
				return
			default:
			}

			fmt.Println(order)

			switch order.Status {
			case model.OrderStatusProcessed:
				c.Repository.SetStatusProcessed(order.Order, order.Accrual)
			case model.OrderStatusInvalid:
				c.Repository.SetStatusInvalid(order.Order)
			case model.OrderStatusRegistered:
			case model.OrderStatusProcessing:
				c.Repository.SetStatusProcessing(order.Order)
			}
		}
	}()
}
