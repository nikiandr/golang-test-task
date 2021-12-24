package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const membersTable = "members"

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewPostgresDBAuthString(auth string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", auth)

	if err != nil {
		return nil, err
	}

	return db, nil
}
