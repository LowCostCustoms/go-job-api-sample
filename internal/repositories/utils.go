package repositories

import (
	"context"
	"fmt"
	"go-scheduler-api/internal/context/vars"
)

type QueryBuilder interface {
	ToSQL() (string, []interface{}, error)
}

func Get[T any](ctx context.Context, queryBuilder QueryBuilder) (*T, error) {
	queryString, queryParams := BuildQuery(queryBuilder)

	var value T
	if err := vars.MustGetQuerier(ctx).Get(&value, queryString, queryParams...); err != nil {
		return nil, err
	}

	return &value, nil
}

func GetCount(ctx context.Context, queryBuilder QueryBuilder) (uint, error) {
	count, err := Get[uint](ctx, queryBuilder)
	if err != nil {
		return 0, err
	} else {
		return *count, nil
	}
}

func Select[T any](ctx context.Context, queryBuilder QueryBuilder) ([]T, error) {
	queryString, queryParams := BuildQuery(queryBuilder)

	var value []T
	if err := vars.MustGetQuerier(ctx).Select(&value, queryString, queryParams...); err != nil {
		return nil, err
	}

	return value, nil
}

func Query(ctx context.Context, queryBuilder QueryBuilder) error {
	queryString, queryParams := BuildQuery(queryBuilder)
	_, err := vars.MustGetQuerier(ctx).Query(queryString, queryParams...)

	return err
}

func BuildQuery(queryBuilder QueryBuilder) (string, []interface{}) {
	if queryString, queryParams, err := queryBuilder.ToSQL(); err != nil {
		panic(fmt.Sprintf("failed to build a query: %v", err))
	} else {
		return queryString, queryParams
	}
}
