package service

import (
	"context"
	"gophermart/internal/logger"
	"gophermart/internal/order/repository"
	"time"
)

type Generator struct {
	doneCh chan struct{}
}

func NewGenerator(doneCh chan struct{}) *Generator {
	return &Generator{
		doneCh: doneCh,
	}
}

// generator возвращает канал с данными
func (g *Generator) Run() chan int {
	// канал, в который будем отправлять данные из слайса
	inputCh := make(chan int)
	sugar := logger.GetLogger()

	// горутина, в которой отправляем в канал  inputCh данные
	go func() {
		// как отправители закрываем канал, когда всё отправим
		defer close(inputCh)

		sugar.Info("Запускаем генератор")
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			// если doneCh закрыт, сразу выходим из горутины
			case <-g.doneCh:
				return
			case <-ticker.C:
				sugar.Info("Прошло 5 секунды! Проверяем заказы")
				orders := getOrders()
				for _, order := range orders {
					inputCh <- order
				}
			}
		}
	}()

	// возвращаем канал для данных
	return inputCh
}

func getOrders() []int {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	orders, err := repository.NewOrderRepository().GetOrdersForAccrual(ctx)
	if err != nil {
		return []int{}
	}
	return orders
}
