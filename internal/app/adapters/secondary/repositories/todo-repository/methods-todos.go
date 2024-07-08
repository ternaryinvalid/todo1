package todo_repository

import (
	"context"

	"github.com/ternaryinvalid/todo1/internal/app/domain/todo"
)

func (repo *TodoRepository) CreateTODO(ctx context.Context, req todo.CreateTaskRequest) (id int, err error) {
	result, err := repo.DB.ExecContext(ctx, "INSERT INTO tasks (title, description, user_id) VALUES (?, ?, ?)", req.Title, req.Description, req.UserID)

	if err != nil {
		return 0, nil
	}

	lastInsertID, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(lastInsertID), nil
}

func (repo *TodoRepository) GetAllTODO(ctx context.Context, userID int) (tasks []todo.Task, err error) {
	rows, err := repo.DB.QueryContext(ctx, "SELECT id, title, description FROM tasks WHERE user_id = ?", userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var taskList []todo.Task
	for rows.Next() {
		var task todo.Task
		err := rows.Scan(&task.TaskID, &task.Title, &task.Description)

		if err != nil {
			return nil, err
		}

		taskList = append(taskList, task)
	}
	return taskList, nil
}

func (repo *TodoRepository) DeleteTODO(ctx context.Context, req todo.DeleteTaskRequest) (err error) {
	_, err = repo.DB.ExecContext(ctx, "DELETE FROM tasks WHERE id = ? AND user_id = ?", req.TaskID, req.UserID)

	if err != nil {
		return err
	}

	return nil

}
