package repositories

import (
	"context"
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

const (
	JobsTableName       = "jobs"
	JobsTableIdColumn   = "id"
	JobsTableNameColumn = "name"
)

type jobRepository struct {
}

func (r *jobRepository) FindById(ctx context.Context, id uuid.UUID) (*Job, error) {
	query := goqu.From(
		JobsTableName,
	).Select(
		JobsTableIdColumn,
		JobsTableNameColumn,
	).Where(
		goqu.Ex{
			JobsTableIdColumn: id,
		},
	)
	job, err := Get[Job](ctx, query)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}

	return job, err
}

func (r *jobRepository) FindAll(ctx context.Context, queryParams JobQueryParams) ([]Job, error) {
	query := filterJobSelectDataset(
		goqu.From(
			JobsTableName,
		).Select(
			JobsTableIdColumn,
			JobsTableNameColumn,
		),
		queryParams,
	).Offset(
		queryParams.Offset,
	)
	if queryParams.Limit != 0 {
		query = query.Limit(queryParams.Limit)
	}

	return Select[Job](ctx, query)
}

func (r *jobRepository) Count(ctx context.Context, queryParams JobQueryParams) (uint, error) {
	query := filterJobSelectDataset(
		goqu.From(JobsTableName).Select(goqu.COUNT("*")),
		queryParams,
	)
	return GetCount(ctx, query)
}

func (r *jobRepository) Create(ctx context.Context, job Job) (*Job, error) {
	job.Id = uuid.New()
	query := goqu.Insert(
		JobsTableName,
	).Cols(
		JobsTableIdColumn,
		JobsTableNameColumn,
	).Vals(
		goqu.Vals{
			job.Id,
			job.Name,
		},
	)
	if err := Query(ctx, query); err != nil {
		return nil, err
	}

	return &job, nil
}

func NewJobRepository() JobRepository {
	return &jobRepository{}
}

func filterJobSelectDataset(ds *goqu.SelectDataset, query JobQueryParams) *goqu.SelectDataset {
	return ds
}
