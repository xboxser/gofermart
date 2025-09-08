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
	if err != nil {

	}

	fmt.Println("user", user)
}
