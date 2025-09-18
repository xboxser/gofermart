package service

import (
	"context"
	"fmt"
	"gophermart/internal/config/db"
	"gophermart/internal/user/model"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Register(apiUser model.ApiUser) (int, error) {
	db := db.GetDB()

	query := `SELECT id, login, password FROM users WHERE login = $1 LIMIT 1`
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	fmt.Println("login", apiUser.Login)
	rows, err := db.Query(ctx, query, string(apiUser.Login))
	if err != nil {
		log.Println("Error query:", err)
		return -1, err
	}
	defer rows.Close()
	// TODO вынести в service отдельное сущность user для более удобного взаимодействия
	var user model.User
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Login, &user.Password)
		if err != nil {
			log.Println("error scan user", err)
			return -1, err
		}
		fmt.Println("user", user)
	}
	fmt.Println("user", user)
	if user.Id != 0 {
		return -1, nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(apiUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
		return -1, err
	}

	userID := -1
	query = `INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id`
	rows, err = db.Query(ctx, query, apiUser.Login, string(hashedPassword))
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
