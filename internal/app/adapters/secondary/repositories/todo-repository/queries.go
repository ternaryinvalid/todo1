package todo_repository

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/ternaryinvalid/todo1/internal/app/domain/todo"
	"github.com/ternaryinvalid/todo1/internal/app/domain/user"
)

func createUserQuery(req user.TodoUserCreateRequest) (string, []interface{}, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query, args, err := psql.Insert("todo.users").
		Columns("username", "password").
		Values(req.Username, req.Password).
		Suffix("RETURNING user_id").
		ToSql()
	if err != nil {
		err = fmt.Errorf("error db [CreateUser]: %v", err)

		return "", nil, err
	}

	return query, args, nil
}

func createTaskQuery(req todo.CreateTaskRequest) (string, []interface{}, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query, args, err := psql.Insert("todo.tasks").
		Columns("user_id", "title", "description", "estimated_date").
		Values(req.UserID, req.Title, req.Description, req.EstimatedDate).
		Suffix("RETURNING task_id").
		ToSql()
	if err != nil {
		err = fmt.Errorf("error db [CreateTask]: %v", err)

		return "", nil, err
	}

	return query, args, nil
}
