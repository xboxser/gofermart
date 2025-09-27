package model

// модель для использования в АПИ
type ApiOrder struct {
	Number  string  `json:"number"`
	Status  string  `json:"status"`
	Accrual float64 `json:"accrual,omitempty"`
	// формат RFC3339
	Uploaded_at string `json:"uploaded_at"`
}
