package middleware

import (
	"gophermart/internal/user/service"
	"log"
	"net/http"
)

// Проверка наличия токена в запросе
func CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получаем значение заголовка
		headerValue := r.Header.Get("Authorization")

		// Здесь можно выполнить нужные действия с полученным значением заголовка
		// Например, логирование или проверку
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

		next.ServeHTTP(w, r)
	})
}
