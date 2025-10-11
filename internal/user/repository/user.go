package repository

import (
	"context"
	"gophermart/internal/config/db"
	"gophermart/internal/logger"
	"gophermart/internal/user/model"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *db.DBPgx
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: db.GetDB(),
	}
}

func (u *UserRepository) RegisterUser(tx pgx.Tx, ctx context.Context, login string, password string) (int, error) {
	sugar := logger.GetLogger()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		sugar.Errorf("Error hashing password: %v", err)
		return 0, err
	}

	userID := 0
	query := `INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id`
	rows, err := tx.Query(ctx, query, login, string(hashedPassword))
	if err != nil {
		sugar.Errorf("Insert error: %v", err)
		return 0, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&userID)
		if err != nil {
			sugar.Errorf("Scan error: %v", err)
			return 0, err
		}
	}

	if userID == 0 {
		return 0, model.ErrRegisterUser
	}
	return userID, nil
}

func (u *UserRepository) GetUserForLogin(ctx context.Context, login string) (model.User, error) {
	db := db.GetDB()
	sugar := logger.GetLogger()
	query := `SELECT id, login, password FROM users WHERE login = $1 LIMIT 1`
	rows, err := db.Query(ctx, query, string(login))
	if err != nil {
		sugar.Errorf("Error query: %v", err)
		return model.User{}, err
	}
	defer rows.Close()

	var user model.User
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Login, &user.Password)
		if err != nil {
			sugar.Error("error scan user %v", err)
			return model.User{}, err
		}
	}

	return user, nil
}

func (u *UserRepository) GetUserForID(ctx context.Context, ID int) (model.User, error) {
	db := db.GetDB()
	sugar := logger.GetLogger()
	query := `SELECT id, login, password FROM users WHERE id = $1 LIMIT 1`
	rows, err := db.Query(ctx, query, ID)
	if err != nil {
		sugar.Errorf("Error query: %v", err)
		return model.User{}, err
	}
	defer rows.Close()

	var user model.User
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Login, &user.Password)
		if err != nil {
			sugar.Errorf("error scan user: %v", err)
			return model.User{}, err
		}
	}

	return user, nil
}
