package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewConfigServer(t *testing.T) {

	oldArgs := os.Args
	oldRunAddr := os.Getenv("RUN_ADDRESS")
	oldDatabaseURI := os.Getenv("DATABASE_URI")
	oldAccrualAddr := os.Getenv("ACCRUAL_SYSTEM_ADDRESS")
	oldAccrualCount := os.Getenv("ACCRUAL_COUNT_CHAN")

	// сохраняем старые значения, чтобы потом восстановить после прохождения тестов
	defer func() {
		os.Args = oldArgs
		os.Setenv("RUN_ADDRESS", oldRunAddr)
		os.Setenv("DATABASE_URI", oldDatabaseURI)
		os.Setenv("ACCRUAL_SYSTEM_ADDRESS", oldAccrualAddr)
		os.Setenv("ACCRUAL_COUNT_CHAN", oldAccrualCount)
	}()

	t.Run("checking the receipt of default parameters", func(t *testing.T) {

		os.Unsetenv("RUN_ADDRESS")
		os.Unsetenv("DATABASE_URI")
		os.Unsetenv("ACCRUAL_SYSTEM_ADDRESS")
		os.Unsetenv("ACCRUAL_COUNT_CHAN")

		// Устанавливаем значения, чтобы не было ошибок у парсера флагов
		os.Args = []string{"program"}

		cfg := NewConfigServer()
		require.NotEmpty(t, cfg.AccrualCountChan)
		require.NotEmpty(t, cfg.AccrualSystemAddress)
		require.NotEmpty(t, cfg.DatabaseURI)
		require.NotEmpty(t, cfg.RunAddress)
	})

	t.Run("checking set env parameters", func(t *testing.T) {

		runAddr := "localhost:8080"
		databaseURI := "postgres://user:password@localhost:5432/database"
		accrualAddr := "localhost:8081"
		os.Setenv("RUN_ADDRESS", runAddr)
		os.Setenv("DATABASE_URI", databaseURI)
		os.Setenv("ACCRUAL_SYSTEM_ADDRESS", accrualAddr)
		os.Setenv("ACCRUAL_COUNT_CHAN", "10")

		// Устанавливаем значения, чтобы не было ошибок у парсера флагов
		os.Args = []string{"program"}

		cfg := NewConfigServer()
		require.Equal(t, cfg.AccrualCountChan, 10)
		require.Equal(t, cfg.AccrualSystemAddress, accrualAddr)
		require.Equal(t, cfg.DatabaseURI, databaseURI)
		require.Equal(t, cfg.RunAddress, runAddr)
	})
}
