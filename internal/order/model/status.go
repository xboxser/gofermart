package model

type StatusOrderType string

const (
	// заказ загружен в систему, но не попал в обработку;
	OrderStatusNew StatusOrderType = "NEW"
	// вознаграждение за заказ рассчитывается;
	OrderStatusProcessing StatusOrderType = "PROCESSING"
	// система расчёта вознаграждений отказала в расчёте;
	OrderStatusInvalid StatusOrderType = "INVALID"
	// данные по заказу проверены и информация о расчёте успешно получена.
	OrderStatusProcessed StatusOrderType = "PROCESSED"
)

type OrderStatus struct {
	List map[StatusOrderType]int
}

func NewOrderStatus() OrderStatus {
	return OrderStatus{
		List: map[StatusOrderType]int{},
	}
}
