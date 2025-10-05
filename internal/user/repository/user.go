package repository

import (
	"context"
	"gophermart/internal/config/db"
	"log"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	ctx context.Context
	db  *db.DBPgx
}

func NewUserRepository(ctx context.Context) *UserRepository {
	return &UserRepository{
		db:  db.GetDB(),
		ctx: ctx,
	}
}

func (u *UserRepository) RegisterUser(tx pgx.Tx, login string, password string) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
		return -1, err
	}

	userID := -1
	query := `INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id`
	rows, err := tx.Query(u.ctx, query, login, string(hashedPassword))
	if err != nil {
		log.Printf("Insert error: %v", err)
		return -1, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&userID)
		if err != nil {
			log.Printf("Scan error: %v", err)
			return -1, err
		}
	}
	return userID, nil

}
