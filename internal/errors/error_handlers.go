package errors

import (
	"context"
	"encoding/json"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go-scheduler-api/internal/context/vars"
	"net/http"
)

var errRouteNotFound = NewNotFoundError("not found")

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewRoutingErrorHandler() runtime.RoutingErrorHandlerFunc {
	return func(
		ctx context.Context,
		mux *runtime.ServeMux,
		marshaler runtime.Marshaler,
		writer http.ResponseWriter,
		request *http.Request,
		code int,
	) {
		runtime.HTTPError(ctx, mux, marshaler, writer, request, errRouteNotFound)
	}
}

func NewErrorHandler() runtime.ErrorHandlerFunc {
	return func(
		ctx context.Context,
		mux *runtime.ServeMux,
		marshaler runtime.Marshaler,
		writer http.ResponseWriter,
		request *http.Request,
		err error,
	) {
		switch err.(type) {
		case *GenericError[NotFoundTag]:
			WriteErrorResponse(ctx, writer, http.StatusNotFound, ErrorResponse{
				Message: err.Error(),
			})
			break
		case *GenericError[BadRequestTag]:
			WriteErrorResponse(ctx, writer, http.StatusBadRequest, ErrorResponse{
				Message: err.Error(),
			})
			break
		default:
			vars.GetLogger(ctx).Errorf("request handler failed: %v", err)

			WriteInternalServerErrorResponse(ctx, writer)
			break
		}
	}
}

func WriteErrorResponse(ctx context.Context, writer http.ResponseWriter, status int, response ErrorResponse) {
	writer.WriteHeader(status)

	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(&response); err != nil {
		logger := vars.GetLogger(ctx)
		logger.Errorf("failed to encode a response body: %v", err)
	}
}

func WriteInternalServerErrorResponse(ctx context.Context, writer http.ResponseWriter) {
	WriteErrorResponse(ctx, writer, http.StatusInternalServerError, ErrorResponse{
		Message: "internal server error",
	})
}
