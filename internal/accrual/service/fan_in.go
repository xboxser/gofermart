package service

import (
	"gophermart/internal/accrual/model"
	"sync"
)

type fanIn struct {
	doneCh chan struct{}
}

func NewFanIn(doneCh chan struct{}) *fanIn {
	return &fanIn{
		doneCh: doneCh,
	}
}

func (f *fanIn) Run(resultChs ...chan model.OrderAccrual) chan model.OrderAccrual {
	resultCh := make(chan model.OrderAccrual)

	var wg sync.WaitGroup

	for _, ch := range resultChs {
		// в горутину передавать переменную цикла нельзя, поэтому делаем так
		chClosure := ch

		wg.Add(1)
		go func() {
			defer wg.Done()
			for data := range chClosure {
				select {
				case <-f.doneCh:
					return
				case resultCh <- data:
				}
			}
		}()
	}

	go func() {
		// ждём завершения всех горутин
		wg.Wait()
		close(resultCh)
	}()

	return resultCh
}
