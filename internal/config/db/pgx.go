package db

// TODO подумать над реализацией интерфейса
import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5"
)

var globalDB *dbPgx
var onceDB sync.Once

type dbPgx struct {
	conn *pgx.Conn
	// TODO перенести PostgresErrorClassifier из проекта метрик
}

func NewDBPgx(ctx context.Context, connStr string) (*dbPgx, error) {
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}

	return &dbPgx{
		conn: conn,
	}, nil
}

func InitDB(ctx context.Context, connStr string) error {
	var initErr error

	onceDB.Do(func() {
		globalDB, initErr = NewDBPgx(ctx, connStr)
	})

	return initErr

}

func GetConn() *pgx.Conn {
	return globalDB.conn
}

func GetDB() *dbPgx {
	return globalDB
}

func (db *dbPgx) Close() error {
	if db.conn != nil {
		return db.conn.Close(context.Background())
	}
	return nil
}
