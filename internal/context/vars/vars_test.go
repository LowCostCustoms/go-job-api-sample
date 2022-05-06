package vars_test

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"go-scheduler-api/internal/context/vars"
	"go-scheduler-api/internal/db"
	"go-scheduler-api/internal/test/logger"
	"testing"
)

func TestGetQuerier(t *testing.T) {
	querier := db.Querier(&sqlx.DB{})
	ctx := vars.WithQuerier(context.Background(), querier)

	assert.Equal(t, vars.GetQuerier(ctx), querier)
}

func TestGetQuerier_NullQuerier(t *testing.T) {
	assert.Nil(t, vars.GetQuerier(context.Background()))
}

func TestGetLogger(t *testing.T) {
	noopLogger := logger.NoopLogger()
	ctx := vars.WithLogger(context.Background(), noopLogger)

	assert.Equal(t, vars.GetLogger(ctx), noopLogger)
}

func TestGetLogger_DefaultLogger(t *testing.T) {
	assert.NotNil(t, vars.GetLogger(context.Background()))
}
