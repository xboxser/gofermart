package handler

import "net/http"

func GetOrders(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("GetOrders"))
}
