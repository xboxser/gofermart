package repository

import (
	"context"
	"gophermart/internal/config/db"
	"gophermart/internal/logger"
	"gophermart/internal/order/model"
)

type StatusRepository struct {
	db *db.DBPgx
}

func NewStatusRepository() *StatusRepository {
	return &StatusRepository{
		db: db.GetDB(),
	}
}

func (s *StatusRepository) GetStatusList(ctx context.Context) (model.OrderStatus, error) {
	sugar := logger.GetLogger()
	status := model.NewOrderStatus()
	status.List = make(map[string]int)
	query := `SELECT id, name FROM statuses`

	rows, err := s.db.Query(ctx, query)
	if err != nil {
		sugar.Errorf("Error query: %v", err)
		return status, err
	}

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			sugar.Errorf("Row scan failed: %v", err)
		}
		status.List[name] = id
	}
	defer rows.Close()

	return status, nil
}
