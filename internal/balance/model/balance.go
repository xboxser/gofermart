package model

type Balance struct {
	// Текущий баланс пользователя
	Current float64 `json:"current"`
	// Сумма потраченных средств
	Withdrawn float64 `json:"withdrawn"`
}
