package main

import (
	"context"
	"fmt"
	"gophermart/internal/config"
	"gophermart/internal/config/db"
	"gophermart/internal/handler"
	"gophermart/internal/validator"
	"log"
)

// TODO добавить логирование
// TODO прикрутить свагер
// TODO Сжатие данных при отправке
// TODO Добавить сессии

func main() {
	config := config.NewConfigServer()
	fmt.Println(config.RunAddress, config.DatabaseURI, config.AccrualSystemAddress)

	ctx := context.Background()
	if err := db.InitDB(ctx, config.DatabaseURI); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.GetDB().Close()

	validator.Init()
	handler := handler.NewServerHandler(config.RunAddress)
	handler.Run()

}
