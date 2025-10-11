package logger

import (
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func Init() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		// вызываем панику, если ошибка
		panic(err)
	}
	defer logger.Sync()
	sugar = logger.Sugar()
}

func GetLogger() *zap.SugaredLogger {
	return sugar
}

func Sync() {
	if sugar != nil {
		sugar.Sync()
	}
}
