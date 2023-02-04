package db

import (
	"database/sql"
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

func (d *DB) Close() error {
	return d.client.Close()
}
