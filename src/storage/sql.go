package storage

import "database/sql"

type SQL struct {
	db *sql.DB
}

func NewSQlDB(client *sql.DB) *SQL {
	return &SQL{db: client}
}
