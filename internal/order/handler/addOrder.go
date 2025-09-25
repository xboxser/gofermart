package handler

import (
	"bytes"
	"gophermart/internal/handler/middleware"
	"gophermart/internal/order/service"
	"gophermart/internal/order/validator"
	"net/http"
)

// 200 — номер заказа уже был загружен этим пользователем;
// 202 — новый номер заказа принят в обработку;
// 400 — неверный формат запроса;
// 401 — пользователь не аутентифицирован;
// 409 — номер заказа уже был загружен другим пользователем;
// 422 — неверный формат номера заказа;
// 500 — внутренняя ошибка сервера.
func AddOrder(res http.ResponseWriter, req *http.Request) {

	userID := req.Context().Value(middleware.UserIDContextKey).(int)
	if userID == 0 {
		http.Error(res, "user not found", http.StatusUnauthorized)
		return
	}

	var buf bytes.Buffer
	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	body := buf.String()
	if body == "" {
		http.Error(res, "body is empty", http.StatusBadRequest)
		return
	}

	IsValidLuhn := validator.IsValidLuhn(body)

	if !IsValidLuhn {
		http.Error(res, "order number is not valid", http.StatusUnprocessableEntity)
		return
	}

	code, err := service.AddOrder(body, userID)
	if err != nil {
		http.Error(res, "error insert order", code)
		return
	}

	res.WriteHeader(code)
}
