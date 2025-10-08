package repository

import (
	"context"
	"gophermart/internal/config/db"
	"gophermart/internal/order/model"
	"log"
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
	status := model.NewOrderStatus()
	status.List = make(map[string]int)
	query := `SELECT id, name FROM statuses`

	rows, err := s.db.Query(ctx, query)
	if err != nil {
		log.Println("Error query:", err)
		return status, err
	}

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal("Row scan failed:", err)
		}
		status.List[name] = id
	}
	defer rows.Close()

	return status, nil
}
