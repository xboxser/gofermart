package client

import (
	"encoding/json"
	"gophermart/internal/accrual/model"
	"net/http"
)

type AccrualClient struct {
	client *http.Client
	url    string
}

func NewAccrualClient(url string) *AccrualClient {
	return &AccrualClient{
		client: &http.Client{},
		url:    url,
	}
}

func (a *AccrualClient) GetOrder(orderNumber string) (model.OrderAccrual, int, error) {
	var order model.OrderAccrual
	res, err := a.client.Get(a.url + "/api/orders/" + orderNumber)
	if err != nil {
		return order, http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return order, res.StatusCode, nil
	}

	err = json.NewDecoder(res.Body).Decode(&order)
	if err != nil {
		return order, res.StatusCode, err
	}
	return order, res.StatusCode, nil
}
