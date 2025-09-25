package service

import (
	"fmt"
	"gophermart/internal/order/repository"
	"net/http"
)

func AddOrder(orderNumber string, userId int) (int, error) {
	orderRepository := repository.NewOrderRepository()
	order, err := orderRepository.GetOrderByNumber(orderNumber)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if order.ID != 0 {
		if userId != order.UserId {
			return http.StatusConflict, nil
		}
		return http.StatusOK, nil
	}
	fmt.Println(userId)
	err = orderRepository.AddOrder(orderNumber, userId)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusAccepted, nil
}
