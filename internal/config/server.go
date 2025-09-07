package config

import (
	"flag"
	"os"

	"github.com/caarlos0/env"
)

type configServer struct {
	RunAddress           string `env:"RUN_ADDRESS"`
	DatabaseURI          string `env:"DATABASE_URI"`
	AccrualSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
}

func NewConfigServer() *configServer {
	var cfg configServer
	_ = env.Parse(&cfg)

	serverFlags := flag.NewFlagSet("server", flag.ExitOnError)
	address := serverFlags.String("a", "localhost:8080", "адрес и порт запуска сервиса")
	databaseURI := serverFlags.String("d", "localhost:5050", "адрес подключения к базе данных")
	accrualAddress := serverFlags.String("r", "localhost:8000", "адрес системы расчёта начислений")
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
	return &cfg
}
