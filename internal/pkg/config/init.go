package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/ternaryinvalid/todo1/internal/app/domain/config"
	"log"
	"strings"
)

func New() (config config.Config) {
	err := cleanenv.ReadConfig("./config.yml", &config)
	if err != nil {
		err = fmt.Errorf(strings.ReplaceAll(err.Error(), ", ", ",\n"))

		log.Fatal(err)
	}

	return
}
