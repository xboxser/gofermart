package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"gophermart/internal/logger"
	"gophermart/internal/user/model"
	"gophermart/internal/user/service"
	"gophermart/internal/user/validator"
	"net/http"
)

// 200 — пользователь успешно зарегистрирован и аутентифицирован;
// 400 — неверный формат запроса;
// 409 — логин уже занят;
// 500 — внутренняя ошибка сервера.
func Register(res http.ResponseWriter, req *http.Request) {
	sugar := logger.GetLogger()

	var buf bytes.Buffer
	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		sugar.Error("error read  body: %v", err)
		return
	}

	var user model.APIUser

	if err = json.Unmarshal(buf.Bytes(), &user); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		sugar.Errorf("error read json: %v", err)
		return
	}

	if err := validator.ValidateModelUserAPI(user); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		sugar.Errorf("Validation failed: %v", err)
	}

	userID, err := service.Register(user)
	if err != nil {
		if errors.Is(err, model.ErrLoginBusy) {
			http.Error(res, "This login is busy", http.StatusConflict)
			sugar.Errorf("This login is busy: %v", err)
			return
		}
		http.Error(res, err.Error(), http.StatusInternalServerError)
		sugar.Errorf("error register: %v", err)
		return
	}

	token, err := service.BuildJWTString(userID)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		sugar.Errorf("error token: %v", err)
		return
	}
	res.Header().Set("Authorization", token)
	res.WriteHeader(http.StatusOK)
}
