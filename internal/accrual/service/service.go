package service

import "fmt"

type AccrualService struct {
	AccrualSystemAddress string
}

func NewAccrualService(AccrualSystemAddress string) *AccrualService {
	return &AccrualService{
		AccrualSystemAddress: AccrualSystemAddress,
	}
}

func (a *AccrualService) Run() {

	go func() {
		// сигнальный канал для завершения горутин
		doneCh := make(chan struct{})
		// закрываем его при завершении программы
		defer close(doneCh)

		// канал с данными
		inputCh := generator(doneCh)

		// запускаем fanOut и получаем результаты обработки данных
		fanOut := NewFanOut(5, a.AccrualSystemAddress, doneCh, inputCh)
		channels := fanOut.Run()

		fanIn := NewFanIn(doneCh)
		resultCh := fanIn.Run(channels...)
		// // а теперь объединяем десять каналов в один
		// addResultCh := fanIn(doneCh, channels...)

		// // передаём тот один канал в следующий этап обработки
		// resultCh := multiply(doneCh, addResultCh)

		// выводим результаты расчетов из канала
		for res := range resultCh {
			fmt.Println("resultCh", res)
		}
	}()

}
