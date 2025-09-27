package service

import (
	"gophermart/internal/order/model"
	"gophermart/internal/order/repository"
	"strconv"
	"time"
)

func GetOrders(userId int) ([]model.APIOrder, error) {
	apiOrders := []model.APIOrder{}
	orders, err := repository.NewOrderRepository().GetOrders(userId)
	if err != nil {
		return []model.APIOrder{}, err
	}

	if len(orders) == 0 {
		return apiOrders, nil
	}

	for _, order := range orders {
		var apiOrder model.APIOrder
		apiOrder.Number = strconv.Itoa(order.Number)
		if order.Accrual != 0 {
			apiOrder.Accrual = order.Accrual
		}

		apiOrder.Status = order.StatusName
		apiOrder.UploadedAt = order.Uploaded_at.Format(time.RFC3339)
		apiOrders = append(apiOrders, apiOrder)
	}

	return apiOrders, nil
}
