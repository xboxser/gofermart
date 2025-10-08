package repository

import (
	"context"
	"gophermart/internal/config/db"
	"gophermart/internal/order/model"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

type OrderRepository struct {
	db *db.DBPgx
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{db: db.GetDB()}
}

// Получаем информацию по заказу по его номеру
func (o *OrderRepository) GetOrderByNumber(ctx context.Context, orderNumber string) (model.Order, error) {
	query := `SELECT orders.id, orders.user_id, orders.accrual, orders.number, statuses.name FROM orders 
	JOIN statuses ON orders.status_id = statuses.id	WHERE number = $1 LIMIT 1`

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

func (o *OrderRepository) AddOrder(ctx context.Context, number string, userID int) error {
	statusRepository := NewStatusRepository()
	statuses, err := statusRepository.GetStatusList(ctx)
	if err != nil {
		return err
	}
	query := `INSERT INTO orders (number, user_id, status_id, accrual) VALUES ($1, $2, $3, $4) RETURNING id`
	_, err = o.db.Exec(ctx, query, number, userID, statuses.List[model.OrderStatusNew], 0)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderRepository) GetOrders(ctx context.Context, userID int) ([]model.Order, error) {
	orders := []model.Order{}
	query := `SELECT orders.id, orders.user_id, orders.accrual, orders.number, orders.uploaded_at, statuses.name FROM orders 
	JOIN statuses ON orders.status_id = statuses.id WHERE user_id = $1`

	rows, err := o.db.Query(ctx, query, userID)
	if err != nil {
		log.Println("Error query:", err)
		return orders, err
	}
	defer rows.Close()

	for rows.Next() {
		var order model.Order
		err := rows.Scan(&order.ID, &order.UserID, &order.Accrual, &order.Number, &order.UploadedAt, &order.StatusName)
		if err != nil {
			log.Println("error scan order", err)
			return orders, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

// получаем список номеров заказа, для проверка сервисом Accrual
func (o *OrderRepository) GetOrdersForAccrual(ctx context.Context) ([]int, error) {
	var orders []int

	query := `SELECT number FROM orders 
	JOIN statuses ON orders.status_id = statuses.id 
	WHERE statuses.name IN ($1, $2)
	`

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
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	query := `UPDATE orders SET status_id = status.id
		FROM status
		WHERE orders.number = $1
  			AND status.name = $2;`
	_, err := o.db.Exec(ctx, query, number, status)
	if err != nil {
		log.Println("Error setStatus query:", err)
		return err
	}
	return nil
}

func (o *OrderRepository) SetStatusProcessed(tx pgx.Tx, ctx context.Context, number string, accrual float64) error {
	query := `UPDATE orders SET status_id = statuses.id, accrual = $1
		FROM statuses
		WHERE orders.number = $2
  			AND statuses.name = $3;`
	_, err := tx.Exec(ctx, query, accrual, number, model.OrderStatusProcessed)
	if err != nil {
		log.Println("Error SetStatusProcessed query:", err)
		return err
	}
	return nil
}
