package middleware

import (
	"context"
	"gophermart/internal/user/handler"
	"gophermart/internal/user/service"
	"log"
	"net/http"
)

// Проверка наличия токена в запросе
func CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerValue := r.Header.Get("Authorization")
		log.Printf("Заголовок %s: %s", "Authorization", headerValue)

		if headerValue == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		userID := service.GetUserID(headerValue)

		if userID == -1 {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		// Добавляем userID в контекст запроса
		// На основе данного поля определяем пользователя в дальнейшем
		ctx := context.WithValue(r.Context(), handler.UserIDContextKey, userID)
		// Передаем запрос с обновленным контекстом дальше
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
