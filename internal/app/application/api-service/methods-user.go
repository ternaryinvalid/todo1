package api_service

import (
	"context"

	"github.com/ternaryinvalid/todo1/internal/app/domain/user"
	"golang.org/x/crypto/bcrypt"
)

func (svc *ApiService) CreateUser(ctx context.Context, req user.TodoUserCreateRequest) (token string, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	req.Password = string(hashedPassword)

	var userID int
	userID, err = svc.userRepository.CreateUser(ctx, req)
	if err != nil {
		return "", err
	}

	token, err = svc.authSvc.GenerateJWT(userID, req.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
