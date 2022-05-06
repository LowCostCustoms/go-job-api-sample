package middleware_test

import (
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"go-scheduler-api/internal/context/vars"
	"go-scheduler-api/internal/middleware"
	http_mocks "go-scheduler-api/internal/test/mocks/http"
	"net/http"
	"testing"
)

type querierContextMatcher struct {
}

func (m *querierContextMatcher) Matches(x interface{}) bool {
	request, _ := x.(*http.Request)
	return request != nil && vars.GetQuerier(request.Context()) != nil
}

func (m *querierContextMatcher) String() string {
	return "has querier context variable"
}

func TestDbContextMiddleware_ServeHTTP(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	nextHandler := http_mocks.NewMockHandler(controller)
	nextHandler.EXPECT().ServeHTTP(gomock.Any(), gomock.Matcher(&querierContextMatcher{})).Times(1)

	responseWriter := http_mocks.NewMockResponseWriter(controller)

	request, _ := http.NewRequest("GET", "http://localhost:3000", nil)

	querier := &sqlx.DB{}

	handler := middleware.NewDbContextMiddleware(querier, nextHandler)

	handler.ServeHTTP(responseWriter, request)
}
