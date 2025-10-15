package handler

import (
	"bytes"
	"encoding/json"
	"gophermart/internal/logger"
	"gophermart/internal/user/model"
	"gophermart/internal/user/service"
	"gophermart/internal/user/validator"
	"net/http"
)

// 200 — пользователь успешно аутентифицирован;
// 400 — неверный формат запроса;
// 401 — неверная пара логин/пароль;
// 500 — внутренняя ошибка сервера.
func Login(res http.ResponseWriter, req *http.Request) {
	sugar := logger.GetLogger()
	var buf bytes.Buffer
	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		sugar.Errorf("error read body: %v", err)
		return
	}

	var user model.APIUser

	if err = json.Unmarshal(buf.Bytes(), &user); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		sugar.Errorf("error read json", err)
		return
	}

	if err := validator.ValidateModelUserAPI(user); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		sugar.Errorf("Validation failed: %v", err)
	}

	userID, err := service.Login(user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		sugar.Errorf("error login: %v", err)
		return
	}

	if userID < 0 {
		http.Error(res, "user not found", http.StatusUnauthorized)
		sugar.Errorf("user not found")
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
