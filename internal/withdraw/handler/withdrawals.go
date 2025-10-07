package handler

import (
	"encoding/json"
	"gophermart/internal/user/handler"
	"gophermart/internal/withdraw/service"
	"net/http"
)

// `200` — успешная обработка запроса.
// `204` — нет ни одного списания.
// `401` — пользователь не авторизован.
// `500` — внутренняя ошибка сервера.
func Withdrawals(res http.ResponseWriter, req *http.Request) {
	userID := handler.GetUserRequest(req)

	service := service.NewWithdrawalsService()
	withdrawals, err := service.Withdrawals(userID)
	if err != nil {
		http.Error(res, "error get withdrawals", http.StatusInternalServerError)
		return
	}

	if len(withdrawals) == 0 {
		res.WriteHeader(http.StatusNoContent)
		return
	}

	resp, err := json.Marshal(withdrawals)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(resp)
}
