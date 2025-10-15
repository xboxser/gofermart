package service

import (
	"fmt"
	"gophermart/internal/accrual/client"
	"gophermart/internal/accrual/model"
	"gophermart/internal/logger"
	"net/http"
	"strconv"
	"time"
)

type fanOut struct {
	countWorkers int
	client       client.AccrualClient
	doneCh       chan struct{}
	inputCh      chan int
	refundCh     chan int
}

func NewFanOut(countWorkers int, accrualSystemAddress string, doneCh chan struct{}, inputCh chan int, refundCh chan int) *fanOut {
	return &fanOut{
		countWorkers: countWorkers,
		client:       *client.NewAccrualClient(accrualSystemAddress),
		doneCh:       doneCh,
		inputCh:      inputCh,
		refundCh:     refundCh,
	}
}

func (f *fanOut) Run() []chan model.OrderAccrual {
	// каналы, в которые отправляются результаты
	channels := make([]chan model.OrderAccrual, f.countWorkers)

	for i := 0; i < f.countWorkers; i++ {
		channels[i] = f.processing()
	}

	return channels
}

func (f *fanOut) processing() chan model.OrderAccrual {
	// канал с результатом
	addRes := make(chan model.OrderAccrual)
	sugar := logger.GetLogger()
	go func() {
		// закрываем канал, когда горутина завершается
		defer close(addRes)

		// берём из канала inputCh номер заказа
		for orderNumber := range f.inputCh {
			order, err := f.sendAccrual(orderNumber)

			sugar.Debugln("result sendAccrual:", orderNumber, order, err)
			select {
			// если канал doneCh закрылся, выходим из горутины
			case <-f.doneCh:
				return
			// если doneCh не закрыт, отправляем результат вычисления в канал результата
			default:
				order.Order = strconv.Itoa(orderNumber)
				if err != nil {
					// возвращаем заказ обратно в очередь
					f.refundCh <- orderNumber
					continue
				}
				addRes <- order
			}
		}
	}()

	return addRes
}

// отправка запроса на сервер Accrual
// `200` — успешная обработка запроса.
// `204` — заказ не зарегистрирован в системе расчета.
// `429` — превышено количество запросов к сервису.

func (f *fanOut) sendAccrual(orderNumber int) (model.OrderAccrual, error) {
	sugar := logger.GetLogger()
	number := strconv.Itoa(orderNumber)
	// Тайминги отправки запроса на сервер
	retryIntervals := []time.Duration{0, 1 * time.Second, 3 * time.Second, 5 * time.Second}

	order, status, err := f.client.GetOrder(number)
	for _, retryInterval := range retryIntervals {
		if retryInterval > 0 {
			sugar.Debugln("retrying time", retryInterval, orderNumber)
			time.Sleep(retryInterval)
			order, status, err = f.client.GetOrder(number)
		}
		// Если сервер вернул ошибку что слишком часто обращаемся  к нему
		// сразу уходим в ретрай
		if status >= 429 {
			continue
		}
		if err == nil {
			break
		}

	}

	if err != nil {
		return order, err
	}

	if status != http.StatusOK {
		return order, fmt.Errorf("status code: %d", status)
	}
	return order, nil

}
