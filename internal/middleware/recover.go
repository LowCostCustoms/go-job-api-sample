package middleware

import (
	"go-scheduler-api/internal/context/vars"
	"go-scheduler-api/internal/errors"
	"net/http"
	"runtime/debug"
)

type RecoverMiddleware struct {
	next http.Handler
}

func (m *RecoverMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			logger := vars.GetLogger(request.Context())
			logger.Errorf("a panic was thrown while handling a request: %s\n%s", err, debug.Stack())

			errors.WriteInternalServerErrorResponse(request.Context(), writer)
		}
	}()

	m.next.ServeHTTP(writer, request)
}

func NewRecoverMiddleware(next http.Handler) http.Handler {
	return &RecoverMiddleware{
		next: next,
	}
}
