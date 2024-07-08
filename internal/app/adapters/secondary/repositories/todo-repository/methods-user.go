package todo_repository

import (
	"context"
	"github.com/ternaryinvalid/todo1/internal/app/domain/user"
)

func (repo *TodoRepository) CreateUser(ctx context.Context, req user.TodoUserCreateRequest) (userID int, err error) {
	query, args, err := createUserQuery(req)
	if err != nil {
		return 0, err
	}

	var dto userIdDTO

	err = repo.DB.GetContext(ctx, &dto, query, args...)
	if err != nil {
		return 0, err
	}

	return dto.UserID, nil
}

type userIdDTO struct {
	UserID int `db:"user_id"`
}
