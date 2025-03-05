package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DB Инициализация pgx
var DB *pgxpool.Pool

// InitDB инициализация нового объекта базы данных
func InitDB(connString string) error {
	var err error
	DB, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}
	return nil
}
