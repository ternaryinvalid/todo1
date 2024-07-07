package api_service

import (
	"context"
	"github.com/ternaryinvalid/todo1/internal/app/domain/todo"
	"github.com/ternaryinvalid/todo1/internal/app/domain/user"
)

type ApiService struct {
	todoRepository todoRepository
	userRepository userRepository
	authSvc        authSvc
}

type authSvc interface {
	GenerateJWT(userID int, username string) (token string, err error)
}

type todoRepository interface {
	CreateTODO(ctx context.Context, req todo.CreateTaskRequest) (id int, err error)
	GetAllTODO(ctx context.Context, userID int) (tasks []todo.Task, err error)
	DeleteTODO(ctx context.Context, req todo.DeleteTaskRequest) (err error)
}

type userRepository interface {
	CreateUser(ctx context.Context, req user.TodoUserCreateRequest) (userID int, err error)
}

func New(
	todoRepository todoRepository,
	userRepository userRepository,
	authSvc authSvc,
) *ApiService {
	return &ApiService{
		todoRepository: todoRepository,
		userRepository: userRepository,
		authSvc:        authSvc,
	}
}
