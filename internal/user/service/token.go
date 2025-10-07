package service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const TokenExp = time.Hour * 3

var jwtSecret = []byte(getJWTSecret())

func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		//TODO проверить как без него будут работать тесты
		secret = "super_secret_code_JWT"
	}
	return secret
}

// Claims — структура утверждений, которая включает стандартные утверждения и
// одно пользовательское UserID
type Claims struct {
	jwt.RegisteredClaims
	UserID int
}

// BuildJWTString создаёт токен и возвращает его в виде строки.
func BuildJWTString(userID int) (string, error) {
	fmt.Println("create token", userID)
	// создаём новый токен с алгоритмом подписи HS256 и утверждениями — Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// когда создан токен
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExp)),
		},
		UserID: userID,
	})

	// создаём строку токена
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	// возвращаем строку токена
	return tokenString, nil
}

func GetUserID(tokenString string) int {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return jwtSecret, nil
		})
	if err != nil {
		return -1
	}

	if !token.Valid {
		return -1
	}

	fmt.Println("Token os valid", claims.UserID)
	return claims.UserID
}
