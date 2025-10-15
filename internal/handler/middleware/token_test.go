package middleware

import (
	"gophermart/internal/user/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock сервиса для тестирования
type mockUserService struct{}

func (m *mockUserService) GetUserID(token string) int {
	if token == "valid_token" {
		return 1
	}
	return -1
}

func (m *mockUserService) GetUserForID(userID int) (model.User, error) {
	if userID == 1 {
		return model.User{ID: 1}, nil
	}
	return model.User{ID: 0}, nil
}

func TestCheckToken(t *testing.T) {
	// Создаем тестовый handler
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Оборачиваем в middleware
	middleware := CheckToken(nextHandler)

	// Проверяем на наличие пустого заголовка Authorization
	t.Run("Empty Authorization header", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		middleware.ServeHTTP(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Errorf("Ожидался статус %d, получен %d", http.StatusUnauthorized, w.Code)
		}
	})

	// Невалидный токен
	t.Run("Invalid token", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", "invalid_token")
		w := httptest.NewRecorder()

		middleware.ServeHTTP(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Errorf("Ожидался статус %d, получен %d", http.StatusUnauthorized, w.Code)
		}
	})
}
