package handler

import "net/http"

func AddOrder(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("AddOrder"))
}
