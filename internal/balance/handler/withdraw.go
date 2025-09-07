package handler

import "net/http"

func Withdraw(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Withdraw"))
}
