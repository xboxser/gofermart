package model

const (
	// заказ загружен в систему, но не попал в обработку;
	OrderStatusNew = "NEW"
	// вознаграждение за заказ рассчитывается;
	OrderStatusProcessing = "PROCESSING"
	// система расчёта вознаграждений отказала в расчёте;
	OrderStatusInvalid = "INVALID"
	// данные по заказу проверены и информация о расчёте успешно получена.
	OrderStatusProcessed = "PROCESSED"
)

type OrderStatus struct {
	List map[string]int
}

func NewOrderStatus() OrderStatus {
	return OrderStatus{
		List: map[string]int{},
	}
}
