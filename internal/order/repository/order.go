package repository

import (
	"context"
	"gophermart/internal/config/db"
	"gophermart/internal/order/model"
	"log"
	"time"
)

type OrderRepository struct {
	db *db.DBPgx
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
		err := rows.Scan(&order.ID, &order.UserID, &order.Accrual, &order.Number, &order.StatusName)
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

func (o *OrderRepository) GetOrders(userId int) ([]model.Order, error) {
	orders := []model.Order{}
	query := `SELECT orders.id, orders.user_id, orders.accrual, orders.number, orders.uploaded_at, statuses.name FROM orders 
	JOIN statuses ON orders.status_id = statuses.id WHERE user_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	rows, err := o.db.Query(ctx, query, userId)
	if err != nil {
		log.Println("Error query:", err)
		return orders, err
	}
	defer rows.Close()

	for rows.Next() {
		var order model.Order
		err := rows.Scan(&order.ID, &order.UserID, &order.Accrual, &order.Number, &order.Uploaded_at, &order.StatusName)
		if err != nil {
			log.Println("error scan order", err)
			return orders, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

// получаем список номеров заказа, для проверка сервисом Accrual
func (o *OrderRepository) GetOrdersForAccrual() ([]int, error) {
	var orders []int

	query := `SELECT number FROM orders 
	JOIN statuses ON orders.status_id = statuses.id 
	WHERE statuses.name IN ($1, $2)
	`
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := o.db.Query(ctx, query, model.OrderStatusProcessing, model.OrderStatusNew)
	if err != nil {
		log.Println("Error query:", err)
		return orders, err
	}
	defer rows.Close()

	for rows.Next() {
		var number int
		err := rows.Scan(&number)
		if err != nil {
			log.Println("error scan order", err)
			return orders, err
		}
		orders = append(orders, number)
	}

	return orders, nil
}

func (o *OrderRepository) SetStatusProcessing(number string) error {
	return o.setStatus(number, model.OrderStatusProcessing)
}

func (o *OrderRepository) SetStatusInvalid(number string) error {
	return o.setStatus(number, model.OrderStatusInvalid)
}

func (o *OrderRepository) setStatus(number string, status string) error {
	query := `UPDATE orders SET status_id = status.id
		FROM status
		WHERE orders.number = $1
  			AND status.name = $2;`
	_, err := o.db.Exec(context.Background(), query, number, status)
	if err != nil {
		log.Println("Error setStatus query:", err)
		return err
	}
	return nil
}

func (o *OrderRepository) SetStatusProcessed(number string, accrual int) error {
	query := `UPDATE orders SET status_id = status.id, accrual = $1
		FROM status
		WHERE orders.number = $2
  			AND status.name = $3;`
	_, err := o.db.Exec(context.Background(), query, accrual, number, model.OrderStatusProcessed)
	if err != nil {
		log.Println("Error SetStatusProcessed query:", err)
		return err
	}
	return nil
}
