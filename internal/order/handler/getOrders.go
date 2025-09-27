package handler

import (
	"encoding/json"
	"gophermart/internal/handler/middleware"
	"gophermart/internal/order/service"
	"net/http"
)

// `200` — успешная обработка запроса.
// `204` — нет данных для ответа.
// `401` — пользователь не авторизован.
// `500` — внутренняя ошибка сервера.
func GetOrders(res http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value(middleware.UserIDContextKey).(int)
	if userID == 0 {
		http.Error(res, "user not found", http.StatusUnauthorized)
		return
	}

	orders, err := service.GetOrders(userID)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(orders) == 0 {
		res.WriteHeader(http.StatusNoContent)
		return
	}

	resp, err := json.Marshal(orders)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(resp)
}
