package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gophermart/internal/user/model"
	"gophermart/internal/user/validator"
	"log"
	"net/http"
)

// TODO 200 — пользователь успешно зарегистрирован и аутентифицирован;
// 400 — неверный формат запроса;
// TODO 409 — логин уже занят;
// TODO 500 — внутренняя ошибка сервера.
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

	if err := validator.ValidateModelUserApi(user); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		log.Println("Validation failed:", err)
	}

	fmt.Println("user", user)
}
