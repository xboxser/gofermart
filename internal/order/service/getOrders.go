package service

import (
	"gophermart/internal/order/model"
	"gophermart/internal/order/repository"
	"strconv"
	"time"
)

func GetOrders(userId int) ([]model.ApiOrder, error) {
	apiOrders := []model.ApiOrder{}
	orders, err := repository.NewOrderRepository().GetOrders(userId)
	if err != nil {
		return []model.ApiOrder{}, err
	}

	if len(orders) == 0 {
		return apiOrders, nil
	}

	for _, order := range orders {
		var apiOrder model.ApiOrder
		apiOrder.Number = strconv.Itoa(order.Number)
		if order.Accrual != 0 {
			apiOrder.Accrual = order.Accrual
		}

		apiOrder.Status = order.StatusName
		apiOrder.Uploaded_at = order.Uploaded_at.Format(time.RFC3339)
		apiOrders = append(apiOrders, apiOrder)
	}

	return apiOrders, nil
}
