package service

import (
	"context"
	"gophermart/internal/config/db"
	"gophermart/internal/user/model"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Register(apiUser model.ApiUser) (int, error) {
	db := db.GetDB()

	user, err := GetUserForLogin(apiUser.Login)

	if user.Id != 0 || err != nil {
		return -1, nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(apiUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
		return -1, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	userID := -1
	query := `INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id`
	rows, err := db.Query(ctx, query, apiUser.Login, string(hashedPassword))
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
