package cmd

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-scheduler-api/internal/container"
)

func NewMigrateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "run database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logrus.StandardLogger()

			migration, err := container.NewMigrate()
			if err != nil {
				logger.Fatalf("failed to configure database migrations: %v", err)
			}

			if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
				logger.Fatalf("failed to apply database migrations: %v", err)
			}

			logger.Info("database migrations applied")
		},
	}
}
