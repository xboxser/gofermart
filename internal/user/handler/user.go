package handler

import (
	"net/http"
)

// избавляемся от проблемы коллизии ключа  в контексте
type ContextKey string

const UserIDContextKey ContextKey = "userID"

func GetUserRequest(req *http.Request) int {
	userID := req.Context().Value(UserIDContextKey).(int)
	return userID
}
