//go:build wireinject
// +build wireinject

package container

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/wire"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"go-scheduler-api/internal/api"
	"go-scheduler-api/internal/config"
	"go-scheduler-api/internal/errors"
	"go-scheduler-api/internal/middleware"
	"go-scheduler-api/internal/repositories"
	"go-scheduler-api/internal/services"
	"net/http"
)

var serviceSet = wire.NewSet(
	services.NewTransactionManager,
	services.NewJobServiceServer,
)

var repositorySet = wire.NewSet(
	repositories.NewJobRepository,
	repositories.NewJobSchedulerRepository,
)

var configSet = wire.NewSet(
	config.GetDatabaseConfig,
	config.GetServerConfig,
)

var containerSet = wire.NewSet(
	repositorySet,
	serviceSet,
	configSet,
	newHttpServer,
	newServeMux,
	newDatabase,
	newMigrate,
)

func NewServer(ctx context.Context, logger logrus.FieldLogger) (*http.Server, error) {
	wire.Build(containerSet)
	return nil, nil
}

func NewMigrate() (*migrate.Migrate, error) {
	wire.Build(containerSet)
	return nil, nil
}

func NewDatabase() (*sqlx.DB, error) {
	wire.Build(containerSet)
	return nil, nil
}

func newDatabase(cfg *config.DatabaseConfig) (*sqlx.DB, error) {
	return sqlx.Open("postgres", cfg.DatabaseUrl)
}

func newHttpServer(
	mux *runtime.ServeMux,
	config *config.ServerConfig,
	db *sqlx.DB,
	logger logrus.FieldLogger,
) (*http.Server, error) {
	handler := http.Handler(mux)
	{
		handler = middleware.NewDbContextMiddleware(db, handler)
		handler = middleware.NewRecoverMiddleware(handler)
		handler = middleware.NewRequestLoggerMiddleware(logger, handler)
	}

	return &http.Server{
		Handler: handler,
		Addr:    fmt.Sprintf(":%d", config.Port),
	}, nil
}

func newServeMux(ctx context.Context, jobServiceServer api.JobServiceServer) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux(
		runtime.WithRoutingErrorHandler(errors.NewRoutingErrorHandler()),
		runtime.WithErrorHandler(errors.NewErrorHandler()),
	)
	if err := api.RegisterJobServiceHandlerServer(ctx, mux, jobServiceServer); err != nil {
		return nil, err
	}

	return mux, nil
}

func newMigrate(cfg *config.DatabaseConfig) (*migrate.Migrate, error) {
	return migrate.New(
		cfg.MigrationsLocation,
		cfg.DatabaseUrl,
	)
}
