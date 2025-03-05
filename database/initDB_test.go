package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitDBGood(t *testing.T) {
	originalPgxPoolNew := pgxpool.New
	defer func() {
		pgxpool.New = originalPgxPoolNew
	}()

	pgxpool.New = func(ctx context.Context, connString string) (*pgxpool.Pool, error) {
		return &pgxpool.Pool{}, nil
	}

	err := InitDB("mock-connection-string")

	assert.NoError(t, err)
}

func TestInitDBFail(t *testing.T) {
	originalPgxPoolNew := pgxpool.New
	defer func() {
		pgxpool.New = originalPgxPoolNew
	}()

	pgxpool.New = func(ctx context.Context, connString string) (*pgxpool.Pool, error) {
		return nil, fmt.Errorf("mock error: unable to connect")
	}

	err := InitDB("mock-connection-string")

	assert.Error(t, err)
	assert.EqualError(t, err, "unable to connect to database: mock error: unable to connect")
}
