package cmd

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-scheduler-api/internal/container"
)

func NewServeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "run web server",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logrus.StandardLogger()
			logger.Formatter = &logrus.TextFormatter{
				TimestampFormat: "2006/01/02 15:04:05.000000",
				FullTimestamp:   true,
			}

			server, err := container.NewServer(context.Background(), logger)
			if err != nil {
				logger.Fatalf("failed to create a server: %v", err)
			}

			logger.Infof("listening on %s", server.Addr)
			if err := server.ListenAndServe(); err != nil {
				logger.Fatalf("failed to start a server: %v", err)
			}
		},
	}
}
