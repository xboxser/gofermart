package service

import (
	"fmt"
	"gophermart/internal/accrual/client"
	"gophermart/internal/accrual/model"
	"net/http"
	"strconv"
)

type fanOut struct {
	countWorkers int
	client       client.AccrualClient
	doneCh       chan struct{}
	inputCh      chan int
}

func NewFanOut(countWorkers int, accrualSystemAddress string, doneCh chan struct{}, inputCh chan int) *fanOut {
	return &fanOut{
		countWorkers: countWorkers,
		client:       *client.NewAccrualClient(accrualSystemAddress),
		doneCh:       doneCh,
		inputCh:      inputCh,
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

	go func() {
		// закрываем канал, когда горутина завершается
		defer close(addRes)

		// берём из канала inputCh номер заказа
		for orderNumber := range f.inputCh {
			fmt.Println("orderNUmber: ", orderNumber)
			order, err := f.sendAccrual(orderNumber)

			fmt.Println("result:", orderNumber, order, err)
			select {
			// если канал doneCh закрылся, выходим из горутины
			case <-f.doneCh:
				return
			// если doneCh не закрыт, отправляем результат вычисления в канал результата
			default:
				if err != nil {
					continue
				}
				addRes <- order
			}
		}
	}()

	return addRes
}

func (f *fanOut) sendAccrual(orderNumber int) (model.OrderAccrual, error) {
	order, err, status := f.client.GetOrder(strconv.Itoa(orderNumber))
	if err != nil {
		return order, err
	}

	if status != http.StatusOK {
		return order, fmt.Errorf("status code: %d", status)
	}
	return order, nil
}
