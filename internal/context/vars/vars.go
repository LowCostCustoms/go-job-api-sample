package vars

import (
	"context"
	"github.com/sirupsen/logrus"
	"go-scheduler-api/internal/db"
)

type querierTag struct{}
type loggerTag struct{}

func WithQuerier(parent context.Context, querier db.Querier) context.Context {
	return context.WithValue(parent, querierTag{}, querier)
}

func GetQuerier(ctx context.Context) db.Querier {
	querier := ctx.Value(querierTag{})
	if querier != nil {
		return querier.(db.Querier)
	}

	return nil
}

func MustGetQuerier(ctx context.Context) db.Querier {
	if querier := GetQuerier(ctx); querier == nil {
		panic("querier interface is missing")
	} else {
		return querier
	}
}

func WithLogger(parent context.Context, logger logrus.FieldLogger) context.Context {
	return context.WithValue(parent, loggerTag{}, logger)
}

func GetLogger(ctx context.Context) logrus.FieldLogger {
	logger := ctx.Value(loggerTag{})
	if logger != nil {
		return logger.(logrus.FieldLogger)
	}

	return logrus.StandardLogger()
}
