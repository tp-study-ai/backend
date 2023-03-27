package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewPostgresqlX() (*sqlx.DB, error) {
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d", "yutfut", "yutfut", "yutfut", "127.0.0.1", "5432")
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1000)
	return db, nil
}
