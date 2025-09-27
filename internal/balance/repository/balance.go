package repository

import (
	"context"
	"gophermart/internal/balance/model"
	"gophermart/internal/config/db"
	"log"
	"time"
)

type BalanceRepository struct {
	db *db.DBPgx
}

func NewBalanceRepository() *BalanceRepository {
	return &BalanceRepository{db: db.GetDB()}
}

func (b *BalanceRepository) GetBalanceUser(userID int) (model.Balance, error) {
	var balance model.Balance
	query := `SELECT current, withdrawn FROM balance 
	WHERE user_id = $1 LIMIT 1`
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	rows, err := b.db.Query(ctx, query, userID)
	if err != nil {
		log.Println("Error query:", err)
		return balance, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&balance.Current, &balance.Withdrawn)
		if err != nil {
			log.Println("error scan balance user", err)
			return balance, err
		}
	}
	return balance, nil
}
