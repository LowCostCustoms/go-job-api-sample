package repositories

import (
	"context"
	"github.com/google/uuid"
)

type Job struct {
	Id   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

type JobSchedule struct {
	Id    uuid.UUID `db:"id"`
	JobId uuid.UUID `db:"job_id"`
	Cron  string    `db:"cron"`
}

type JobQueryParams struct {
	Offset uint
	Limit  uint
}

type JobRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (*Job, error)
	FindAll(ctx context.Context, query JobQueryParams) ([]Job, error)
	Count(ctx context.Context, query JobQueryParams) (uint, error)
	Create(ctx context.Context, job Job) (*Job, error)
}

type JobScheduleRepository interface {
	FindByIds(ctx context.Context, id ...uuid.UUID) ([]JobSchedule, error)
	FindByJobId(ctx context.Context, jobId uuid.UUID) ([]JobSchedule, error)
	CreateMany(ctx context.Context, schedules ...JobSchedule) ([]JobSchedule, error)
}
