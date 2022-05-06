package db

import (
	"database/sql"
)

// Querier defines a set of methods for querying SQL data.
type Querier interface {
	Query(query string, queryParams ...interface{}) (*sql.Rows, error)
	Get(dest interface{}, query string, queryParams ...interface{}) error
	Select(dest interface{}, query string, queryParams ...interface{}) error
}

// Transaction defines a database transaction.
type Transaction interface {
	Commit() error
	Rollback() error
}
