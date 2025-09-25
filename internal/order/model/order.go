package model

type Order struct {
	ID         int     `db:"id"`
	Number     int     `db:"number"`
	Accrual    float64 `db:"accrual"`
	StatusName string  `db:"status_id"`
	UserId     int     `db:"user_id"`
}
