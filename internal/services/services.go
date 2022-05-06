package services

import (
	"context"
	"go-scheduler-api/internal/db"
)

type TransactionManager interface {
	Begin(ctx context.Context) (db.Querier, db.Transaction, error)
}
