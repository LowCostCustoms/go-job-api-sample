package repositories

import (
	"context"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

const (
	JobSchedulesTableName        = "job_schedules"
	JobSchedulesTableIdColumn    = "id"
	JobSchedulesTableJobIdColumn = "job_id"
	JobSchedulesTableCronColumn  = "cron"
)

type jobSchedulerRepository struct {
}

func (r *jobSchedulerRepository) FindByIds(ctx context.Context, id ...uuid.UUID) ([]JobSchedule, error) {
	query := goqu.From(
		JobSchedulesTableName,
	).Select(
		JobSchedulesTableIdColumn,
		JobSchedulesTableJobIdColumn,
		JobSchedulesTableCronColumn,
	).Where(
		goqu.Ex{
			JobSchedulesTableIdColumn: id,
		},
	)
	return Select[JobSchedule](ctx, query)
}

func (r *jobSchedulerRepository) FindByJobId(ctx context.Context, jobId uuid.UUID) ([]JobSchedule, error) {
	query := goqu.From(
		JobSchedulesTableName,
	).Select(
		JobSchedulesTableIdColumn,
		JobSchedulesTableJobIdColumn,
		JobSchedulesTableCronColumn,
	).Where(
		goqu.Ex{
			JobSchedulesTableJobIdColumn: jobId,
		},
	)
	return Select[JobSchedule](ctx, query)
}

func (r *jobSchedulerRepository) CreateMany(ctx context.Context, schedules ...JobSchedule) ([]JobSchedule, error) {
	if len(schedules) == 0 {
		return nil, nil
	}

	rows := make([][]any, len(schedules))
	for idx, schedule := range schedules {
		jobScheduleId := uuid.New()
		schedules[idx] = JobSchedule{
			Id:    jobScheduleId,
			JobId: schedule.JobId,
			Cron:  schedule.Cron,
		}
		rows[idx] = goqu.Vals{
			jobScheduleId,
			schedule.JobId,
			schedule.Cron,
		}
	}

	query := goqu.Insert(
		JobSchedulesTableName,
	).Cols(
		JobSchedulesTableIdColumn,
		JobSchedulesTableJobIdColumn,
		JobSchedulesTableCronColumn,
	).Vals(
		rows...,
	)
	if err := Query(ctx, query); err != nil {
		return nil, err
	}

	return schedules, nil
}

func NewJobSchedulerRepository() JobScheduleRepository {
	return &jobSchedulerRepository{}
}
