package storage

import "database/sql"

type Postgres struct {
	db *sql.DB
}

func NewPostgresDB(client *sql.DB) *Postgres {
	return &Postgres{db: client}
}
