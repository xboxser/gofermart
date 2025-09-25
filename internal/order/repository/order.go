package repository

import (
	"context"
	"gophermart/internal/config/db"
	"gophermart/internal/order/model"
	"log"
	"time"
)

type OrderRepository struct {
	db *db.DbPgx
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{db: db.GetDB()}
}

func (o *OrderRepository) GetOrderByNumber(orderNumber string) (model.Order, error) {

	query := `SELECT orders.id, orders.user_id, orders.accrual, orders.number, statuses.name FROM orders 
	JOIN statuses ON orders.status_id = statuses.id	WHERE number = $1 LIMIT 1`
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	rows, err := o.db.Query(ctx, query, orderNumber)
	if err != nil {
		log.Println("Error query:", err)
		return model.Order{}, err
	}
	defer rows.Close()

	var order model.Order
	if rows.Next() {
		err := rows.Scan(&order.ID, &order.UserId, &order.Accrual, &order.Number, &order.StatusName)
		if err != nil {
			log.Println("error scan order", err)
			return model.Order{}, err
		}
	}
	return order, nil
}

func (o *OrderRepository) AddOrder(number string, userId int) error {
	statusRepository := NewStatusRepository()
	statuses, err := statusRepository.GetStatusList()
	if err != nil {
		return err
	}
	query := `INSERT INTO orders (number, user_id, status_id, accrual) VALUES ($1, $2, $3, $4) RETURNING id`
	_, err = o.db.Exec(context.Background(), query, number, userId, statuses.List[model.OrderStatusNew], 0)
	if err != nil {
		return err
	}
	return nil
}
