package model

type OrderAccrualStatus string

const (
	// заказ зарегистрирован, но не начисление не рассчитано;
	OrderStatusRegistered OrderAccrualStatus = "REGISTERED"
	// заказ не принят к расчёту, и вознаграждение не будет начислено;
	OrderStatusInvalid OrderAccrualStatus = "INVALID"
	// расчёт начисления в процессе;
	OrderStatusProcessing OrderAccrualStatus = "PROCESSING"
	// расчёт начисления окончен
	OrderStatusProcessed OrderAccrualStatus = "PROCESSED"
)

type OrderAccrual struct {
	Order   string             `json:"order"`
	Status  OrderAccrualStatus `json:"status"`
	Accrual float64            `json:"accrual,omitempty"`
}
