package validator

import (
	"errors"
	"gophermart/internal/user/model"
	projectValidator "gophermart/internal/validator"

	"github.com/go-playground/validator/v10"
)

// Проверяем обязательные поля
// Скрываем ошибки формата "Key: 'ApiUser.Login' Error:Field validation for 'Login' failed on the 'required' tag"
func ValidateModelUserAPI(user model.APIUser) error {
	err := projectValidator.Validate.Struct(user)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "Login":
					return errors.New("Login is required")
				case "Password":
					return errors.New("Password is required")
				}
			}
		} else {
			return errors.New("Invalid input data")
		}

	}
	return nil
}
