package service

import (
	"context"
	"gophermart/internal/order/model"
	"gophermart/internal/order/repository"
	"time"
)

func AddOrder(orderNumber string, userID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	orderRepository := repository.NewOrderRepository()
	order, err := orderRepository.GetOrderByNumber(ctx, orderNumber)
	if err != nil {
		return err
	}

	if order.ID != 0 {
		if userID != order.UserID {
			return model.Error409
		}
		return nil
	}

	err = orderRepository.AddOrder(ctx, orderNumber, userID)
	if err != nil {
		return err
	}
	return model.Error202
}
