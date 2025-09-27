package handler

import (
	"net/http"
)

const UserIDContextKey = "userID"

func GetUserRequest(req *http.Request) int {
	userID := req.Context().Value(UserIDContextKey).(int)
	return userID
}
