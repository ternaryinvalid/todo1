package api_service

import (
	"context"

	"github.com/ternaryinvalid/todo1/internal/app/domain/todo"
)

func (svc *ApiService) CreateTODO(ctx context.Context, req todo.CreateTaskRequest) (id int, err error) {
	id, err = svc.todoRepository.CreateTODO(ctx, req)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (svc *ApiService) DeleteTODO(ctx context.Context, req todo.DeleteTaskRequest) (err error) {
	err = svc.todoRepository.DeleteTODO(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (svc *ApiService) GetAllTODO(ctx context.Context, userID int) (tasks []todo.Task, err error) {
	tasks, err = svc.todoRepository.GetAllTODO(ctx, userID)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
