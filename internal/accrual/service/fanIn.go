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

// // fanIn объединяет несколько каналов resultChs в один.
// func fanIn(doneCh chan struct{}, resultChs ...chan int) chan int {
// 	// конечный выходной канал в который отправляем данные из всех каналов из слайса, назовём его результирующим
// 	finalCh := make(chan int)

// 	// понадобится для ожидания всех горутин
// 	var wg sync.WaitGroup

// 	// перебираем все входящие каналы
// 	for _, ch := range resultChs {
// 		// в горутину передавать переменную цикла нельзя, поэтому делаем так
// 		chClosure := ch

// 		// инкрементируем счётчик горутин, которые нужно подождать
// 		wg.Add(1)

// 		go func() {
// 			// откладываем сообщение о том, что горутина завершилась
// 			defer wg.Done()

// 			// получаем данные из канала
// 			for data := range chClosure {
// 				select {
// 				// выходим из горутины, если канал закрылся
// 				case <-doneCh:
// 					return
// 				// если не закрылся, отправляем данные в конечный выходной канал
// 				case finalCh <- data:
// 				}
// 			}
// 		}()
// 	}

// 	go func() {
// 		// ждём завершения всех горутин
// 		wg.Wait()
// 		// когда все горутины завершились, закрываем результирующий канал
// 		close(finalCh)
// 	}()

// 	// возвращаем результирующий канал
// 	return finalCh
// }
