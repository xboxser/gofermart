package model

type Withdraw struct {
	OrderID string  `json:"order"`
	Sum     float64 `json:"sum"`
}
