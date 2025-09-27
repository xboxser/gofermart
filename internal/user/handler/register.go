package handler

import (
	"bytes"
	"encoding/json"
	"gophermart/internal/user/model"
	"gophermart/internal/user/service"
	"gophermart/internal/user/validator"
	"log"
	"net/http"
)

// 200 — пользователь успешно зарегистрирован и аутентифицирован;
// 400 — неверный формат запроса;
// 409 — логин уже занят;
// 500 — внутренняя ошибка сервера.
func Register(res http.ResponseWriter, req *http.Request) {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		log.Println("error read  body", err)
		return
	}

	var user model.ApiUser

	if err = json.Unmarshal(buf.Bytes(), &user); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		log.Println("error read json", err)
		return
	}

	if err := validator.ValidateModelUserAPI(user); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		log.Println("Validation failed:", err)
	}

	userID, err := service.Register(user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		log.Println("error register", err)
		return
	}

	if userID == -1 {
		http.Error(res, "This login is busy", http.StatusConflict)
		return
	}

	token, err := service.BuildJWTString(userID)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		log.Println("error token", err)
		return
	}
	res.Header().Set("Authorization", token)
	res.WriteHeader(http.StatusOK)
}
