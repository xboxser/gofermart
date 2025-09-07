package handler

import "net/http"

func Withdrawals(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Withdrawals"))
}
