package infra

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewPostgresDB() (*sql.DB, error) {
	// Параметры подключения (такие же как в docker-compose.yml)
	connStr := "user=auth_user password=auth_pass dbname=auth_db host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %v", err)
	}

	// Проверяем, что подключение работает
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db ping failed: %v", err)
	}

	return db, nil
}

func InitDB(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT UNIQUE,
			avatar_url TEXT
		)
	`)
	return err
}
