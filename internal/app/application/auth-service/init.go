package auth_service

import (
	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	jwtSecret string
}

type Claims struct {
	Username string `json:"username"`
	UserID   int    `json:"user_id"`
	jwt.StandardClaims
}

func New(jwtSecret string) *AuthService {
	return &AuthService{
		jwtSecret: jwtSecret,
	}
}
