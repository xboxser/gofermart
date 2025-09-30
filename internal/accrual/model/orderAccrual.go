package model

const (
	// заказ зарегистрирован, но не начисление не рассчитано;
	OrderStatusRegistered = "REGISTERED"
	// заказ не принят к расчёту, и вознаграждение не будет начислено;
	OrderStatusInvalid = "INVALID"
	// расчёт начисления в процессе;
	OrderStatusProcessing = "PROCESSING"
	// расчёт начисления окончен
	OrderStatusProcessed = "PROCESSED"
)

type OrderAccrual struct {
	Order   string `json:"order"`
	Status  string `json:"status"`
	Accrual int    `json:"accrual,omitempty"`
}
