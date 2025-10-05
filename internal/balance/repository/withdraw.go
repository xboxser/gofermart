package repository

import (
	"context"
	"gophermart/internal/config/db"

	"github.com/jackc/pgx/v5"
)

type WithdrawRepository struct {
	db *db.DBPgx
}

func NewWithdrawRepository(db *db.DBPgx) *WithdrawRepository {
	return &WithdrawRepository{db: db}
}

func (w *WithdrawRepository) AddWithdraw(tx pgx.Tx, ctx context.Context, userID int, amount float64, orderNumber string) error {
	_, err := tx.Exec(ctx, "INSERT INTO withdrawal (user_id, sum, order_id) VALUES ($1, $2, $3)", userID, amount, orderNumber)
	return err
}
