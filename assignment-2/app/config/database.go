package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	DB_HOST = "localhost"
	DB_PORT = "55000"
	DB_USER = "postgres"
	DB_PASS = "postgrespw"
	DB_NAME = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(10 * time.Second)
	db.SetConnMaxLifetime(10 * time.Second)

	return db, nil
}
