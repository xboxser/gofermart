package handler

import "net/http"

func GetBalance(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Get Balance"))
}
