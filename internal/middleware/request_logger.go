package middleware

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go-scheduler-api/internal/context/vars"
	"net/http"
	"time"
)

type RequestLoggerMiddleware struct {
	logger logrus.FieldLogger
	next   http.Handler
}

func (m *RequestLoggerMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	requestId := uuid.New()
	logger := m.logger.WithField("requestId", requestId)
	nextRequest := request.WithContext(vars.WithLogger(request.Context(), logger))
	writerInterceptor := NewResponseWriterInterceptor(writer)

	start := time.Now()
	m.next.ServeHTTP(writerInterceptor, nextRequest)

	logger.Infof(
		"request %s %s finished with HTTP %d, took %v",
		request.Method,
		request.URL.Path,
		writerInterceptor.Status(),
		time.Since(start),
	)
}

func NewRequestLoggerMiddleware(logger logrus.FieldLogger, next http.Handler) http.Handler {
	return &RequestLoggerMiddleware{
		logger: logger,
		next:   next,
	}
}

type ResponseWriterInterceptor struct {
	next       http.ResponseWriter
	statusCode int
}

func (w *ResponseWriterInterceptor) Header() http.Header {
	return w.next.Header()
}

func (w *ResponseWriterInterceptor) Write(bytes []byte) (int, error) {
	return w.next.Write(bytes)
}

func (w *ResponseWriterInterceptor) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.next.WriteHeader(statusCode)
}

func (w *ResponseWriterInterceptor) Status() int {
	return w.statusCode
}

func NewResponseWriterInterceptor(next http.ResponseWriter) *ResponseWriterInterceptor {
	return &ResponseWriterInterceptor{
		next:       next,
		statusCode: http.StatusOK,
	}
}
