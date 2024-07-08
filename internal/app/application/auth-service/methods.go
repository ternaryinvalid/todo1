package auth_service

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// AuthMiddleware функция для проверки JWT
func (svc *AuthService) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := extractTokenFromHeader(authHeader)
		if tokenStr == "" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		claims := &Claims{}

		token, err := jwt.ParseWithClaims(
			tokenStr,
			claims,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(svc.jwtSecret), nil
			},
		)
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)

			return
		}

		// Сохраняем ID пользователя в контексте запроса
		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// GenerateJWT генерирует JWT токен для пользователя
func (svc *AuthService) GenerateJWT(userID int, username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(svc.jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func extractTokenFromHeader(authHeader string) string {
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}
	return parts[1]
}
