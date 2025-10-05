package model

import "time"

type ApiWithdraw struct {
	OrderID    string    `json:"order"`
	Sum        float64   `json:"sum"`
	UploadedAt time.Time `json:"processed_at"`
}
