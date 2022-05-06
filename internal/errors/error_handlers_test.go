package errors_test

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go-scheduler-api/internal/errors"
	http_mocks "go-scheduler-api/internal/test/mocks/http"
	"net/http"
	"testing"
)

type writeInternalServerErrorTestCase struct {
	name  string
	error error
}

type errorHandlerTestCase struct {
	name           string
	error          error
	expectedStatus int
}

func TestWriteInternalServerErrorResponse(t *testing.T) {
	testCases := []writeInternalServerErrorTestCase{
		{
			name:  "NoError",
			error: nil,
		},
		{
			name:  "Error",
			error: fmt.Errorf("error"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			writer := http_mocks.NewMockResponseWriter(controller)
			writer.EXPECT().WriteHeader(http.StatusInternalServerError).Times(1)
			writer.EXPECT().Write(gomock.Any()).DoAndReturn(func(data []byte) (int, error) {
				return len(data), testCase.error
			})

			errors.WriteInternalServerErrorResponse(context.Background(), writer)
		})
	}
}

func TestErrorHandler(t *testing.T) {
	testCases := []errorHandlerTestCase{
		{
			name:           "NotFoundError",
			error:          errors.NewNotFoundError("not found"),
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "BadRequestError",
			error:          errors.NewBadRequestError("bad request"),
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Other",
			error:          fmt.Errorf("unknown error"),
			expectedStatus: http.StatusInternalServerError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			writer := http_mocks.NewMockResponseWriter(controller)
			writer.EXPECT().Write(gomock.Any())
			writer.EXPECT().WriteHeader(testCase.expectedStatus).Times(1)

			errorHandler := errors.NewErrorHandler()

			errorHandler(context.Background(), nil, nil, writer, nil, testCase.error)
		})
	}
}

func TestRoutingErrorHandler(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	routingErrorHandler := errors.NewRoutingErrorHandler()
	mux := runtime.NewServeMux(
		runtime.WithErrorHandler(errors.NewErrorHandler()),
		runtime.WithRoutingErrorHandler(routingErrorHandler),
	)

	writer := http_mocks.NewMockResponseWriter(controller)
	writer.EXPECT().Write(gomock.Any())
	writer.EXPECT().WriteHeader(http.StatusNotFound).Times(1)

	routingErrorHandler(context.Background(), mux, nil, writer, nil, 100)
}
