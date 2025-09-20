package service

import (
	"context"
	"gophermart/internal/config/db"
	"gophermart/internal/user/model"
	"log"
	"time"
)

func GetUserForLogin(login string) (model.User, error) {
	db := db.GetDB()

	query := `SELECT id, login, password FROM users WHERE login = $1 LIMIT 1`
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	rows, err := db.Query(ctx, query, string(login))
	if err != nil {
		log.Println("Error query:", err)
		return model.User{}, err
	}
	defer rows.Close()

	var user model.User
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Login, &user.Password)
		if err != nil {
			log.Println("error scan user", err)
			return model.User{}, err
		}
	}

	return user, nil
}
