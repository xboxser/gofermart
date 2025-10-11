package service

type AccrualService struct {
	AccrualSystemAddress string
	AccrualCountChan     int
}

func NewAccrualService(AccrualSystemAddress string, AccrualCountChan int) *AccrualService {
	return &AccrualService{
		AccrualSystemAddress: AccrualSystemAddress,
		AccrualCountChan:     AccrualCountChan,
	}
}

func (a *AccrualService) Run() {

	go func() {
		// сигнальный канал для завершения горутин
		doneCh := make(chan struct{})

		// закрываем его при завершении программы
		// defer close(doneCh)

		// канал с данными
		generator := NewGenerator(doneCh)
		generatorCh := generator.Run()

		// запускаем fanOut и получаем результаты обработки данных
		fanOut := NewFanOut(a.AccrualCountChan, a.AccrualSystemAddress, doneCh, generatorCh)
		channels := fanOut.Run()

		fanIn := NewFanIn(doneCh)
		resultCh := fanIn.Run(channels...)

		consumer := NewConsumer(resultCh, doneCh)
		consumer.Run()

	}()

}
