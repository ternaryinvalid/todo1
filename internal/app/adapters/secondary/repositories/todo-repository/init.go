package todo_repository

import (
	"database/sql"
	"fmt"
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
