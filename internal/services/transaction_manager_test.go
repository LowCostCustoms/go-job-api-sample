package services_test

import (
	"context"
	"database/sql"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"go-scheduler-api/internal/context/vars"
	"go-scheduler-api/internal/services"
	db_mocks "go-scheduler-api/internal/test/mocks/db"
	services_mocks "go-scheduler-api/internal/test/mocks/services"
	"testing"
)

func TestTransactionManager_Begin(t *testing.T) {
	db, _ := sqlx.Open("sqlite3", ":memory:")
	defer db.Close()

	ctx := vars.WithQuerier(context.Background(), db)
	transactionManager := services.NewTransactionManager()

	querier, newTx, err := transactionManager.Begin(ctx)

	assert.NotNil(t, querier)
	assert.NotNil(t, newTx)
	assert.Nil(t, err)
	assert.IsType(t, &sqlx.Tx{}, newTx)
}

func TestTransactionManager_Begin_WithinTransactionScope(t *testing.T) {
	db, _ := sqlx.Open("sqlite3", ":memory:")
	defer db.Close()

	tx, _ := db.Beginx()
	defer tx.Rollback()

	ctx := vars.WithQuerier(context.Background(), tx)
	transactionManager := services.NewTransactionManager()

	querier, newTx, err := transactionManager.Begin(ctx)

	assert.NotNil(t, querier)
	assert.NotNil(t, newTx)
	assert.Nil(t, err)
	assert.IsType(t, services.NewNoopTransaction(), newTx)
}

func TestNoopTransaction_Commit(t *testing.T) {
	assert.Nil(t, services.NewNoopTransaction().Commit())
}

func TestNoopTransaction_Rollback(t *testing.T) {
	assert.Nil(t, services.NewNoopTransaction().Rollback())
}

func TestWithTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	expectedValue := 100

	transaction := db_mocks.NewMockTransaction(controller)
	transaction.EXPECT().Commit().Times(1).Return(nil)
	transaction.EXPECT().Rollback().Times(1).Return(sql.ErrTxDone)

	querier := db_mocks.NewMockQuerier(controller)
	querier.EXPECT().Query("query").Times(1)

	transactionManager := services_mocks.NewMockTransactionManager(controller)
	transactionManager.EXPECT().Begin(gomock.Any()).Times(1).Return(querier, transaction, nil)

	value, err := services.WithTransaction(
		context.Background(),
		transactionManager,
		func(ctx context.Context) (*int, error) {
			_, _ = vars.GetQuerier(ctx).Query("query")
			return &expectedValue, nil
		},
	)

	assert.Nil(t, err)
	assert.Equal(t, value, &expectedValue)
}

func TestWithTransaction_ForwardError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	expectedError := errors.New("some error")

	transaction := db_mocks.NewMockTransaction(controller)
	transaction.EXPECT().Rollback().Times(1).Return(sql.ErrTxDone)

	querier := db_mocks.NewMockQuerier(controller)

	transactionManager := services_mocks.NewMockTransactionManager(controller)
	transactionManager.EXPECT().Begin(gomock.Any()).Times(1).Return(querier, transaction, nil)

	value, err := services.WithTransaction(
		context.Background(),
		transactionManager,
		func(ctx context.Context) (*int, error) {
			return nil, expectedError
		},
	)

	assert.Nil(t, value)
	assert.NotNil(t, err)
	assert.Same(t, err, expectedError)
}
