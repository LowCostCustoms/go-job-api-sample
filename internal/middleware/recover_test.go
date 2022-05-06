package middleware_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"go-scheduler-api/internal/context/vars"
	"go-scheduler-api/internal/middleware"
	"go-scheduler-api/internal/test/logger"
	http_mocks "go-scheduler-api/internal/test/mocks/http"
	"net/http"
	"testing"
)

func TestRecoverMiddleware_ServeHTTP(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	panicHandler := http_mocks.NewMockHandler(controller)
	panicHandler.EXPECT().ServeHTTP(
		gomock.Any(),
		gomock.Any(),
	).Times(1).Do(
		func(w http.ResponseWriter, r *http.Request) {
			panic("error")
		},
	)

	handler := middleware.NewRecoverMiddleware(panicHandler)

	request, _ := http.NewRequestWithContext(
		vars.WithLogger(
			context.Background(),
			logger.NoopLogger(),
		),
		"GET",
		"http://localhost:3000",
		nil,
	)

	responseWriter := http_mocks.NewMockResponseWriter(controller)
	responseWriter.EXPECT().WriteHeader(http.StatusInternalServerError).Times(1)
	responseWriter.EXPECT().Write(gomock.Any()).Times(1)

	handler.ServeHTTP(responseWriter, request)
}
