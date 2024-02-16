package persistence

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/caarlos0/env"
)

type config struct {
	User      string `env:"POSTGRES_USER" envDefault:"user"`
	Password  string `env:"POSTGRES_PASSWORD" envDefault:"password"`
	Port      string `env:"POSTGRES_PORT" envDefault:"5432"`
	Host      string `env:"POSTGRES_HOST" envDefault:""`
	DefaultDB string `env:"POSTGRES_DEFAULTDB" envDefault:"stori"`
}

func OpenDataBaseConnection() (*sql.DB, error) {

	cfg := config{}

	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	cnxString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable options='-csearch_path=stori'", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DefaultDB)
	dbConnection, err := sql.Open("postgres", cnxString)

	if err != nil {
		return nil, err
	}

	if err := dbConnection.Ping(); err != nil {
		return nil, err
	}

	return dbConnection, nil
}
