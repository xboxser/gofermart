package handler

import "net/http"

func Register(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Register"))
}
