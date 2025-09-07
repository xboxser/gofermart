package main

import (
	"fmt"
	balance "gophermart/internal/balance/handler"
	order "gophermart/internal/order/handler"
	user "gophermart/internal/user/handler"
	withdraw "gophermart/internal/withdraw/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// TODO добавить логирование
// TODO прикрутить свагер
// TODO Сжатие данных при отправке

func main() {

	fmt.Print("Hello World!")

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
		Addr:    "localhost:8080",
		Handler: r,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
