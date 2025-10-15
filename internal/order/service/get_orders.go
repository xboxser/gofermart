package service

import (
	"context"
	"gophermart/internal/order/model"
	"gophermart/internal/order/repository"
	"strconv"
	"time"
)

func GetOrders(userID int) ([]model.APIOrder, error) {
	apiOrders := []model.APIOrder{}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	orders, err := repository.NewOrderRepository().GetOrders(ctx, userID)
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
		apiOrder.UploadedAt = order.UploadedAt.Format(time.RFC3339)
		apiOrders = append(apiOrders, apiOrder)
	}

	return apiOrders, nil
}
