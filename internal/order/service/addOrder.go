package service

import (
	"context"
	"gophermart/internal/order/repository"
	"net/http"
	"time"
)

func AddOrder(orderNumber string, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	orderRepository := repository.NewOrderRepository()
	order, err := orderRepository.GetOrderByNumber(ctx, orderNumber)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if order.ID != 0 {
		if userID != order.UserID {
			return http.StatusConflict, nil
		}
		return http.StatusOK, nil
	}

	err = orderRepository.AddOrder(ctx, orderNumber, userID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusAccepted, nil
}
