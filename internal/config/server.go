package config

import (
	"flag"
	"os"

	"github.com/caarlos0/env"
)

type ConfigServer struct {
	RunAddress           string `env:"RUN_ADDRESS"`
	DatabaseURI          string `env:"DATABASE_URI"`
	AccrualSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	AccrualCountChan     int    `env:"ACCRUAL_COUNT_CHAN"`
}

func NewConfigServer() *ConfigServer {
	var cfg ConfigServer
	_ = env.Parse(&cfg)

	serverFlags := flag.NewFlagSet("server", flag.ExitOnError)
	address := serverFlags.String("a", "localhost:8008", "адрес и порт запуска сервиса")
	//
	databaseURI := serverFlags.String("d", "postgres://gofermart_user:qwerty!@3@localhost:5432/gofermart_db?sslmode=disable", "адрес подключения к базе данных")
	accrualAddress := serverFlags.String("r", "http://localhost:8080", "адрес системы расчёта начислений")
	accrualCountChan := serverFlags.Int("c", 5, "количество потоков выгружаемых заказов в Accrual")
	serverFlags.Parse(os.Args[1:])

	if cfg.RunAddress == "" {
		cfg.RunAddress = *address
	}

	if cfg.DatabaseURI == "" {
		cfg.DatabaseURI = *databaseURI
	}

	if cfg.AccrualSystemAddress == "" {
		cfg.AccrualSystemAddress = *accrualAddress
	}

	if cfg.AccrualCountChan == 0 {
		cfg.AccrualCountChan = *accrualCountChan
	}
	return &cfg
}
