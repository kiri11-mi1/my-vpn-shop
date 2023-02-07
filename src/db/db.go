package db

import (
	"context"
	"database/sql"
	"time"
)

type DB struct {
	client *sql.DB
}

func Connect(driver, connStr string) (*DB, error) {
	db, err := sql.Open(driver, connStr)
	if err != nil {
		return nil, err
	}
	return &DB{client: db}, nil
}

func (d *DB) Client() *sql.DB {
	return d.client
}

func (d *DB) CreateTables() error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	if _, err := d.client.ExecContext(ctx, QUERY); err != nil {
		return err
	}
	return nil
}

func (d *DB) Close() error {
	return d.client.Close()
}
