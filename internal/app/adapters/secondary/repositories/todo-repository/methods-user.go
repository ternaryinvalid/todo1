package todo_repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"github.com/ternaryinvalid/todo1/internal/app/domain/user"
)

const (
	AlreadyExist = "23505" // duplicate key value violates unique constraint "idx_reports_source"
)

func (repo *TodoRepository) CreateUser(ctx context.Context, req user.TodoUserCreateRequest) (userID int, err error) {
	query, args, err := createUserQuery(req)
	if err != nil {
		return 0, err
	}

	var dto userIdDTO

	err = repo.DB.GetContext(ctx, &dto, query, args...)
	if err != nil {
		var pqError *pq.Error
		if errors.As(err, &pqError) {
			if pqError.Code == AlreadyExist {
				err = fmt.Errorf("error: user already exists")

				return 0, err
			}
		}

		return 0, err
	}

	return dto.UserID, nil
}

type userIdDTO struct {
	UserID int `db:"user_id"`
}
