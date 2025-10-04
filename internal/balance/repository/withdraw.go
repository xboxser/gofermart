package repository

import "gophermart/internal/config/db"

type WithdrawRepository struct {
	db *db.DBPgx
}

func NewWithdrawRepository(db *db.DBPgx) *WithdrawRepository {
	return &WithdrawRepository{db: db}
}
