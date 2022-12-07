package storage

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type GoJWTDb interface{}

type Db struct {
	*sql.DB
}

func NewDB() (*Db, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("GO_JWT_TEMPLATE_DB_HOST"),
		os.Getenv("GO_JWT_TEMPLATE_DB_PORT"),
		os.Getenv("GO_JWT_TEMPLATE_DB_USER"),
		os.Getenv("GO_JWT_TEMPLATE_DB_PASSWORD"),
		os.Getenv("GO_JWT_TEMPLATE_DB_NAME"),
		os.Getenv("GO_JWT_TEMPLATE_SSL_MODE"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Db{db}, nil
}
