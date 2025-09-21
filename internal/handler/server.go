package handler

import (
	balance "gophermart/internal/balance/handler"
	"gophermart/internal/handler/middleware"
	order "gophermart/internal/order/handler"
	user "gophermart/internal/user/handler"
	withdraw "gophermart/internal/withdraw/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type serverHandler struct {
	addr string
}

func NewServerHandler(addr string) *serverHandler {
	return &serverHandler{
		addr: addr,
	}
}

func (h *serverHandler) Run() {
	r := chi.NewRouter()
	r.Route("/api/user", func(r chi.Router) {
		r.Post("/register", user.Register)
		r.Post("/login", user.Login)

		r.Route("/orders", func(r chi.Router) {
			r.Use(middleware.CheckToken)
			r.Post("/", order.AddOrder)
			r.Get("/", order.GetOrders)
		})
		r.Route("/balance", func(r chi.Router) {
			r.Use(middleware.CheckToken)
			r.Get("/", balance.GetBalance)
			r.Post("/withdraw", balance.Withdraw)
		})

		r.Route("/withdrawals", func(r chi.Router) {
			r.Use(middleware.CheckToken)
			r.Get("/", withdraw.Withdrawals)
		})

	})
	server := &http.Server{
		Addr:    h.addr,
		Handler: r,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
