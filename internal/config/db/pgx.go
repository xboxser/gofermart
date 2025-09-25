package db

// TODO подумать над реализацией интерфейса
import (
	"context"
	"embed"
	"errors"
	"fmt"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var globalDB *DbPgx
var onceDB sync.Once

type DbPgx struct {
	conn *pgx.Conn
	// TODO перенести PostgresErrorClassifier из проекта метрик
}

// gofermart_user qwerty!@3

func NewDBPgx(ctx context.Context, connStr string) (*DbPgx, error) {
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}

	return &DbPgx{
		conn: conn,
	}, nil
}

func InitDB(ctx context.Context, connStr string) error {
	var initErr error

	onceDB.Do(func() {
		globalDB, initErr = NewDBPgx(ctx, connStr)
		if initErr != nil {
			return
		}
		initErr = runMigrations(connStr)
	})
	fmt.Printf("initDB: %v", initErr)
	return initErr

}

// TODO скорее всего использоваться не будет и нужно удалить
func GetConn() *pgx.Conn {
	return globalDB.conn
}

func GetDB() *DbPgx {
	return globalDB
}

func (db *DbPgx) Close() error {
	if db.conn != nil {
		return db.conn.Close(context.Background())
	}
	return nil
}

//go:embed migrations/*.sql
var migrationsDir embed.FS

func runMigrations(dsn string) error {
	d, err := iofs.New(migrationsDir, "migrations")
	if err != nil {
		return fmt.Errorf("failed to return an iofs driver: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, dsn)
	if err != nil {
		return fmt.Errorf("failed to get a new migrate instance: %w", err)
	}

	// if err := m.Down(); err != nil {
	// 	if !errors.Is(err, migrate.ErrNoChange) {
	// 		return fmt.Errorf("failed to rollback migration: %w", err)
	// 	}
	// }
	if err := m.Up(); err != nil {

		// TODO для переключения версии миграции
		// if _, dirty := err.(migrate.ErrDirty); dirty {
		// 	// Принудительно устанавливаем версию
		// 	m.Force(3) // или нужный номер версии
		// } else {
		// 	log.Fatal(err)
		// }

		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("failed to apply migrations to the DB: %w", err)
		}
	}
	return nil
}

func (db *DbPgx) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return db.conn.Query(ctx, sql, args...)
}

func (db *DbPgx) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	return db.conn.Exec(ctx, sql, args...)
}
