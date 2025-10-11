package handler

import (
	"encoding/json"
	"gophermart/internal/balance/service"
	"gophermart/internal/user/handler"
	"net/http"
)

// `200` — успешная обработка запроса.
// `401` — пользователь не авторизован.
// `500` — внутренняя ошибка сервера.
func GetBalance(res http.ResponseWriter, req *http.Request) {
	userID := handler.GetUserRequest(req)

	service := service.NewGetBalance(userID)
	balance, err := service.GetBalanceUser()

	if err != nil {
		http.Error(res, "error get balance", http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(balance)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(resp)
}
