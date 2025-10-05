package handler

import (
	"bytes"
	"encoding/json"
	"gophermart/internal/balance/model"
	"gophermart/internal/balance/service"
	"gophermart/internal/order/validator"
	"gophermart/internal/user/handler"

	"log"
	"net/http"
)

// 200 — успешная обработка запроса;
// 401 — пользователь не авторизован;
// 402 — на счету недостаточно средств;
// 422 — неверный номер заказа;
// 500 — внутренняя ошибка сервера.
func Withdraw(res http.ResponseWriter, req *http.Request) {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		log.Println("error read  body", err)
		return
	}

	var withdrawModel model.Withdraw

	if err = json.Unmarshal(buf.Bytes(), &withdrawModel); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		log.Println("error read json", err)
		return
	}

	IsValidLuhn := validator.IsValidLuhn(withdrawModel.OrderID)

	if !IsValidLuhn {
		http.Error(res, "order number is not valid", http.StatusUnprocessableEntity)
		return
	}

	userID := handler.GetUserRequest(req)
	if userID == 0 {
		http.Error(res, "user not found", http.StatusUnauthorized)
		return
	}

	withdrawService := service.NewWithdraw()
	status, err := withdrawService.Withdraw(userID, withdrawModel)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(status)
}
