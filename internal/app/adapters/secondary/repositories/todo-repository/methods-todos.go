package todo_repository

import (
	"context"
	"log"

	"github.com/ternaryinvalid/todo1/internal/app/domain/todo"
)

func (repo *TodoRepository) CreateTODO(ctx context.Context, req todo.CreateTaskRequest) (id int, err error) {
	query, args, err := createTaskQuery(req)
	if err != nil {
		return 0, nil
	}

	var taskIdDto taskIDDTO

	err = repo.DB.GetContext(ctx, &taskIdDto, query, args...)
	if err != nil {
		return 0, err
	}

	return taskIdDto.TaskId, nil
}

func (repo *TodoRepository) GetAllTODO(ctx context.Context, userID int) (tasks []todo.Task, err error) {
	rows, err := repo.DB.QueryContext(ctx, "SELECT t.task_id, t.title, t.description, t.estimated_date, t.done FROM todo.tasks t WHERE t.user_id = $1", userID)

	if err != nil {
		return nil, err
	}

	defer func() {
		dbErr := rows.Close()
		if dbErr != nil {
			log.Println(dbErr)
		}
	}()

	var taskList []todo.Task
	for rows.Next() {
		var task todo.Task
		err := rows.Scan(&task.TaskID, &task.Title, &task.Description, &task.EstimatedDate, &task.Done)

		if err != nil {
			return nil, err
		}

		taskList = append(taskList, task)
	}
	return taskList, nil
}

func (repo *TodoRepository) DeleteTODO(ctx context.Context, req todo.DeleteTaskRequest) (err error) {
	_, err = repo.DB.ExecContext(ctx, "DELETE FROM todo.tasks t WHERE t.task_id = $1 AND user_id = $2", req.TaskID, req.UserID)

	if err != nil {
		return err
	}

	return nil

}

func (repo *TodoRepository) Done(ctx context.Context, req todo.DoneTaskRequest) (err error) {
	_, err = repo.DB.ExecContext(ctx, "UPDATE todo.tasks SET done = true WHERE task_id = $1 AND user_id = $2;", req.TaskID, req.UserID)
	if err != nil {
		return err
	}

	return nil
}

type taskIDDTO struct {
	TaskId int `db:"task_id"`
}
