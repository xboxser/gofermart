package service

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
		// defer close(doneCh)

		// канал с данными
		inputCh := generator(doneCh)

		// запускаем fanOut и получаем результаты обработки данных
		fanOut := NewFanOut(5, a.AccrualSystemAddress, doneCh, inputCh)
		channels := fanOut.Run()

		fanIn := NewFanIn(doneCh)
		resultCh := fanIn.Run(channels...)

		consumer := NewConsumer(resultCh, doneCh)
		consumer.Run()

	}()

}
