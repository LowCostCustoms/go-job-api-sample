package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-scheduler-api/internal/context/vars"
	"go-scheduler-api/internal/db"
)

var (
	defaultNoopTransaction = NewNoopTransaction()
	errMissingDbContext    = errors.New("database context is missing")
)

type noopTransaction struct {
}

func (t *noopTransaction) Commit() error {
	return nil
}

func (t *noopTransaction) Rollback() error {
	return nil
}

func NewNoopTransaction() db.Transaction {
	return &noopTransaction{}
}

type transactionManager struct {
}

func (m *transactionManager) Begin(ctx context.Context) (db.Querier, db.Transaction, error) {
	querier := vars.GetQuerier(ctx)
	if querier == nil {
		return nil, nil, errMissingDbContext
	}

	connection, _ := querier.(*sqlx.DB)
	if connection != nil {
		tx, err := connection.Beginx()
		return tx, tx, err
	}

	return querier, defaultNoopTransaction, nil
}

func NewTransactionManager() TransactionManager {
	return &transactionManager{}
}

func WithTransaction[T any](
	ctx context.Context,
	mgr TransactionManager,
	handler func(ctx context.Context) (*T, error),
) (*T, error) {
	querier, tx, err := mgr.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
			logger := vars.GetLogger(ctx)
			logger.Errorf("failed to rollback a transaction: %v", err)
		}
	}()

	reply, err := handler(vars.WithQuerier(ctx, querier))
	if err == nil {
		if commitErr := tx.Commit(); commitErr != nil {
			return nil, fmt.Errorf("failed to commit a transaction: %v", commitErr)
		}
	}

	return reply, err
}
