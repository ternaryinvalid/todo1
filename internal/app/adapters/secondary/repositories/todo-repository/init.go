package todo_repository

import (
	"fmt"
	sql "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ternaryinvalid/todo1/internal/app/domain/config"
	"github.com/ternaryinvalid/todo1/internal/pkg/repohelpers"
	"log"
	"os"
)

type TodoRepository struct {
	config config.Database
	DB     *sql.DB
}

func New(cfg config.Database) *TodoRepository {
	currentHostString := fmt.Sprintf("DB host: [%s:%s].", cfg.Host, cfg.Port)

	log.Println(currentHostString + " Подключение...")

	connectionString := repohelpers.GetConnectionString(cfg.Type, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := sql.Open(cfg.Type, connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println(err.Error())

		os.Exit(1)
	}

	log.Println(currentHostString + " Подключено!")
	log.Println(currentHostString + " Подключено!")

	return &TodoRepository{
		config: cfg,
		DB:     db,
	}
}

func (repo *TodoRepository) Init() {
	createTablesQuery := `
CREATE SCHEMA IF NOT EXISTS todo;

ALTER SCHEMA todo OWNER TO user_admin;

CREATE TABLE IF NOT EXISTS todo.users (
    user_id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY,
    username text UNIQUE NOT NULL,
    password text NOT NULL,
    registered_at timestamp without time zone DEFAULT now() NOT NULL,
    PRIMARY KEY (user_id)
);

ALTER TABLE todo.users OWNER TO user_admin;

CREATE TABLE IF NOT EXISTS todo.tasks (
    task_id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY,
    user_id BIGINT NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    estimated_date text NOT NULL,
    done BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (task_id),  -- Define primary key constraint here
    FOREIGN KEY (user_id) REFERENCES todo.users(user_id) MATCH FULL
);

ALTER TABLE todo.tasks OWNER TO user_admin;
`

	_, err := repo.DB.Exec(createTablesQuery)
	if err != nil {
		log.Fatal(err)
	}
}
