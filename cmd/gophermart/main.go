package main

import (
	"context"
	"gophermart/internal/accrual/service"
	"gophermart/internal/config"
	"gophermart/internal/config/db"
	"gophermart/internal/handler"
	"gophermart/internal/logger"
	"gophermart/internal/validator"
)

// TODO Сжатие данных при отправке

func main() {
	config := config.NewConfigServer()

	logger.Init()
	defer logger.Sync()
	sugar := logger.GetLogger()

	sugar.Infof("Starting server at %s", config.RunAddress)

	ctx := context.Background()
	if err := db.InitDB(ctx, config.DatabaseURI); err != nil {
		sugar.Fatalf("Failed to initialize database: %v", err)
	}

	defer func() {
		if err := db.GetDB().Close(); err != nil {
			sugar.Fatalf("Error close db: %v", err)
		}
	}()

	validator.Init()

	handler := handler.NewServerHandler(config.RunAddress)

	accrualService := service.NewAccrualService(config.AccrualSystemAddress, config.AccrualCountChan)
	accrualService.Run()
	handler.Run()
}
