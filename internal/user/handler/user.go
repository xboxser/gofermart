package handler

import (
	"net/http"
)

// избавляемся от проблемы коллизии ключа  в контексте
type ContextKey string

const UserIDContextKey ContextKey = "userID"

// получение ID пользователя из контекста
// корректность ID проверяется в middleware
func GetUserRequest(req *http.Request) int {
	userID := req.Context().Value(UserIDContextKey).(int)
	return userID
}
