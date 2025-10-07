package repository

import (
	"context"
	"gophermart/internal/config/db"
	"gophermart/internal/user/model"
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

func (u *UserRepository) GetUserForLogin(login string) (model.User, error) {
	db := db.GetDB()

	query := `SELECT id, login, password FROM users WHERE login = $1 LIMIT 1`
	rows, err := db.Query(u.ctx, query, string(login))
	if err != nil {
		log.Println("Error query:", err)
		return model.User{}, err
	}
	defer rows.Close()

	var user model.User
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Login, &user.Password)
		if err != nil {
			log.Println("error scan user", err)
			return model.User{}, err
		}
	}

	return user, nil
}

func (u *UserRepository) GetUserForId(ID int) (model.User, error) {
	db := db.GetDB()

	query := `SELECT id, login, password FROM users WHERE id = $1 LIMIT 1`
	rows, err := db.Query(u.ctx, query, ID)
	if err != nil {
		log.Println("Error query:", err)
		return model.User{}, err
	}
	defer rows.Close()

	var user model.User
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Login, &user.Password)
		if err != nil {
			log.Println("error scan user", err)
			return model.User{}, err
		}
	}

	return user, nil
}
