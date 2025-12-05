package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	Login(email, password string) (*LoginResponse, error)
}

type LoginResponse struct {
	Token     string
	Email     string
	ExpiresIn int64
}

type authService struct{}

func NewAuthService() AuthService {
	return &authService{}
}

// To make the test simple, we use hardcoded users
var users = map[string]string{
	"admin@email.com": "password123",
	"user@email.com":  "userpass",
}

func (service *authService) Login(email, password string) (*LoginResponse, error) {
	storedPassword, exists := users[email]
	if !exists || storedPassword != password {
		return nil, errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(15 * time.Minute).Unix(),
		"iat":   time.Now().Unix(),
	})

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, errors.New("JWT_SECRET environment variable not set")
	}

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token:     tokenString,
		Email:     email,
		ExpiresIn: 15 * 60, // 15 minutes in seconds
	}, nil
}
