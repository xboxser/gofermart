package model

const (
	OrderStatusNew        = "NEW"
	OrderStatusProcessing = "PROCESSING"
	OrderStatusInvalid    = "INVALID"
	OrderStatusProcessed  = "PROCESSED"
)

type OrderStatus struct {
	List map[string]int
}

func NewOrderStatus() OrderStatus {
	return OrderStatus{
		List: map[string]int{},
	}
}
