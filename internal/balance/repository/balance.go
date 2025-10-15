package repository

import (
	"context"
	"gophermart/internal/balance/model"
	"gophermart/internal/config/db"
	"gophermart/internal/logger"

	"github.com/jackc/pgx/v5"
)

type BalanceRepository struct {
	db *db.DBPgx
}

func NewBalanceRepository() *BalanceRepository {
	return &BalanceRepository{db: db.GetDB()}
}

func (b *BalanceRepository) GetBalanceUser(ctx context.Context, userID int) (model.Balance, error) {
	sugar := logger.GetLogger()
	var balance model.Balance
	//Устанавливаем отрицательное значение withdrawn, чтобы в случае ошибки в дальнейшем можно было понять, что balance не найден
	balance.Withdrawn = -1
	query := `SELECT current, withdrawn FROM balance 
	WHERE user_id = $1 LIMIT 1`
	rows, err := b.db.Query(ctx, query, userID)
	if err != nil {
		sugar.Errorf("Error query: %v", err)
		return balance, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&balance.Current, &balance.Withdrawn)
		if err != nil {
			sugar.Errorf("error scan balance user: %v", err)
			return balance, err
		}
	}
	return balance, nil
}

func (b *BalanceRepository) CreateBalanceUser(tx pgx.Tx, ctx context.Context, userID int) error {
	query := `INSERT INTO balance (user_id, current, withdrawn) VALUES ($1, 0, 0)`
	_, err := tx.Exec(ctx, query, userID)
	if err != nil {
		return err
	}
	return err
}

func (b *BalanceRepository) AddBalanceUser(tx pgx.Tx, ctx context.Context, userID int, amount float64) error {
	query := `UPDATE balance SET current = current + $1 WHERE user_id = $2`
	_, err := tx.Exec(ctx, query, amount, userID)
	if err != nil {
		return err
	}
	return err
}

// Вывод средств с баланса
func (b *BalanceRepository) WithdrawBalanceUser(tx pgx.Tx, ctx context.Context, userID int, amount float64) error {
	query := `UPDATE balance SET current = current - $1, withdrawn = withdrawn + $1 WHERE user_id = $2`
	_, err := tx.Exec(ctx, query, amount, userID)
	if err != nil {
		return err
	}
	return err
}
