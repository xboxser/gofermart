package middleware

import (
	"context"
	"fmt"
	"gophermart/internal/user/handler"
	"gophermart/internal/user/service"
	"net/http"
)

// Проверка наличия токена в запросе
func CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerValue := r.Header.Get("Authorization")

		if headerValue == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		fmt.Println(headerValue)
		// Проверяем что передали числовое значение
		userID := service.GetUserID(headerValue)
		if userID == -1 {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Ищем пользователя в базе данных
		user, err := service.GetUserForID(userID)
		if user.ID < 1 || err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Добавляем userID в контекст запроса
		// На основе данного поля определяем пользователя в дальнейшем
		ctx := context.WithValue(r.Context(), handler.UserIDContextKey, user.ID)
		// Передаем запрос с обновленным контекстом дальше
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
