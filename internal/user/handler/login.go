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

// 200 — пользователь успешно аутентифицирован;
// 400 — неверный формат запроса;
// 401 — неверная пара логин/пароль;
// 500 — внутренняя ошибка сервера.
func Login(res http.ResponseWriter, req *http.Request) {
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

	if err := validator.ValidateModelUserApi(user); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		log.Println("Validation failed:", err)
	}

	userID, err := service.Login(user)
	if err != nil {
		if userID > 0 {
			http.Error(res, "user not found", http.StatusUnauthorized)
			return
		}
		http.Error(res, err.Error(), http.StatusInternalServerError)
		log.Println("error login", err)
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
