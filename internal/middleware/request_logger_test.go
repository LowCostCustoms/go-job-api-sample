package middleware_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-scheduler-api/internal/context/vars"
	"go-scheduler-api/internal/middleware"
	"go-scheduler-api/internal/test/logger"
	http_mocks "go-scheduler-api/internal/test/mocks/http"
	"net/http"
	"testing"
)

type hasLoggerContextMatcher struct {
}

func (m *hasLoggerContextMatcher) Matches(x interface{}) bool {
	request, _ := x.(*http.Request)
	return request != nil && vars.GetLogger(request.Context()) != nil
}

func (m *hasLoggerContextMatcher) String() string {
	return "has logger context variable"
}

func TestRequestLoggerMiddleware_ServeHTTP(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	nextHandler := http_mocks.NewMockHandler(controller)
	nextHandler.EXPECT().ServeHTTP(gomock.Any(), gomock.Matcher(&hasLoggerContextMatcher{})).Times(1)

	request, _ := http.NewRequestWithContext(
		context.Background(),
		"GET",
		"http://localhost:3000",
		nil,
	)

	handler := middleware.NewRequestLoggerMiddleware(logger.NoopLogger(), nextHandler)

	handler.ServeHTTP(nil, request)
}

func TestResponseWriterInterceptor_Header(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	nextWriter := http_mocks.NewMockResponseWriter(controller)
	nextWriter.EXPECT().Header().Times(1)

	writer := middleware.NewResponseWriterInterceptor(nextWriter)

	writer.Header()
}

func TestResponseWriterInterceptor_Write(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	data := []byte("stuff")
	expectedCount := 10
	expectedError := errors.New("some-error")

	nextWriter := http_mocks.NewMockResponseWriter(controller)
	nextWriter.EXPECT().Write(data).Times(1).Return(expectedCount, expectedError)

	writer := middleware.NewResponseWriterInterceptor(nextWriter)

	count, err := writer.Write(data)

	assert.Equal(t, count, expectedCount)
	assert.Equal(t, err, expectedError)
}

func TestResponseWriterInterceptor_WriteHeader(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	status := http.StatusBadRequest

	nextWriter := http_mocks.NewMockResponseWriter(controller)
	nextWriter.EXPECT().WriteHeader(status).Times(1)

	writer := middleware.NewResponseWriterInterceptor(nextWriter)

	writer.WriteHeader(status)

	assert.Equal(t, writer.Status(), status)
}
