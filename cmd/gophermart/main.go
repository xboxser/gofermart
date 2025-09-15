package main

import (
	"context"
	"fmt"
	balance "gophermart/internal/balance/handler"
	"gophermart/internal/config"
	"gophermart/internal/config/db"
	order "gophermart/internal/order/handler"
	user "gophermart/internal/user/handler"
	"gophermart/internal/validator"
	withdraw "gophermart/internal/withdraw/handler"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// TODO добавить логирование
// TODO прикрутить свагер
// TODO Сжатие данных при отправке

func main() {
	config := config.NewConfigServer()
	fmt.Println(config.RunAddress, config.DatabaseURI, config.AccrualSystemAddress)

	ctx := context.Background()
	if err := db.InitDB(ctx, config.DatabaseURI); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.GetDB().Close()

	validator.Init()
	r := chi.NewRouter()
	r.Route("/api/user", func(r chi.Router) {
		r.Post("/register", user.Register)
		r.Post("/login", user.Login)

		r.Route("/orders", func(r chi.Router) {
			r.Post("/", order.AddOrder)
			r.Get("/", order.GetOrders)
		})
		r.Route("/balance", func(r chi.Router) {
			r.Get("/", balance.GetBalance)
			r.Post("/withdraw", balance.Withdraw)
		})

		r.Get("/withdrawals", withdraw.Withdrawals)
	})
	server := &http.Server{
		Addr:    config.RunAddress,
		Handler: r,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
