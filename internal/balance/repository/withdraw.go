package repository

import (
	"context"
	"gophermart/internal/config/db"
	"gophermart/internal/withdraw/model"

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

func (w *WithdrawRepository) GetWithdraws(userID int) ([]model.APIWithdraw, error) {
	var withdraws []model.APIWithdraw
	rows, err := w.db.Query(context.Background(), "SELECT order_id, sum, update_at FROM withdrawal WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var withdraw model.APIWithdraw
		err := rows.Scan(&withdraw.OrderID, &withdraw.Sum, &withdraw.UploadedAt)
		if err != nil {
			return nil, err
		}
		withdraws = append(withdraws, withdraw)
	}
	return withdraws, nil

}
