package todo_repository

import (
	"context"
	"github.com/ternaryinvalid/todo1/internal/app/domain/todo"
)

func (repo *TodoRepository) CreateTODO(ctx context.Context, req todo.CreateTaskRequest) (id int, err error) {
	//TODO implement me
	panic("implement me")
}

func (repo *TodoRepository) GetAllTODO(ctx context.Context, userID int) (tasks []todo.Task, err error) {
	//TODO implement me
	panic("implement me")
}

func (repo *TodoRepository) DeleteTODO(ctx context.Context, req todo.DeleteTaskRequest) (err error) {
	//TODO implement me
	panic("implement me")
}
