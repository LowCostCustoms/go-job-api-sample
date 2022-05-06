package middleware

import (
	"go-scheduler-api/internal/context/vars"
	"go-scheduler-api/internal/db"
	"net/http"
)

type dbContextMiddleware struct {
	querier db.Querier
	next    http.Handler
}

func (m *dbContextMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	m.next.ServeHTTP(
		writer,
		request.WithContext(
			vars.WithQuerier(
				request.Context(),
				m.querier,
			),
		),
	)
}

func NewDbContextMiddleware(querier db.Querier, next http.Handler) http.Handler {
	return &dbContextMiddleware{
		querier: querier,
		next:    next,
	}
}
