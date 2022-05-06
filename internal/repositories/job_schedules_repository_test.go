package repositories_test

import (
	"context"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"go-scheduler-api/internal/context/vars"
	"go-scheduler-api/internal/repositories"
	"go-scheduler-api/internal/test/db"
	"testing"
)

func TestJobSchedulerRepository_FindByIds(t *testing.T) {
	db.WithTestDatabase(func(tx *sqlx.Tx) {
		job := createJob(tx)
		schedules := []repositories.JobSchedule{
			createJobSchedule(tx, job.Id),
			createJobSchedule(tx, job.Id),
			createJobSchedule(tx, job.Id),
		}
		repository := repositories.NewJobSchedulerRepository()

		foundSchedules, err := repository.FindByIds(createContext(tx), schedules[0].Id, schedules[1].Id)

		assert.Nil(t, err)
		assert.Len(t, foundSchedules, 2)
		assert.Contains(t, foundSchedules, schedules[0])
		assert.Contains(t, foundSchedules, schedules[1])
	})
}

func TestJobSchedulerRepository_CreateMany(t *testing.T) {
	db.WithTestDatabase(func(tx *sqlx.Tx) {
		job := createJob(tx)
		schedules := []repositories.JobSchedule{
			{
				JobId: job.Id,
				Cron:  "some-cron",
			},
			{
				JobId: job.Id,
				Cron:  "other-cron",
			},
		}
		repository := repositories.NewJobSchedulerRepository()
		ctx := vars.WithQuerier(context.Background(), tx)
		createdSchedules, err := repository.CreateMany(ctx, schedules...)

		assert.Nil(t, err)
		assert.Len(t, createdSchedules, len(schedules))

		foundSchedules, err := repository.FindByIds(ctx, createdSchedules[0].Id, createdSchedules[1].Id)

		assert.Nil(t, err)
		assert.ElementsMatch(t, createdSchedules, foundSchedules)
	})
}

func TestJobSchedulerRepository_FindByJobId(t *testing.T) {
	db.WithTestDatabase(func(tx *sqlx.Tx) {
		job := createJob(tx)
		otherJob := createJob(tx)
		schedules := []repositories.JobSchedule{
			createJobSchedule(tx, job.Id),
			createJobSchedule(tx, job.Id),
			createJobSchedule(tx, otherJob.Id),
			createJobSchedule(tx, otherJob.Id),
		}
		repository := repositories.NewJobSchedulerRepository()

		foundSchedules, err := repository.FindByJobId(createContext(tx), job.Id)

		assert.Nil(t, err)
		assert.Len(t, foundSchedules, 2)
		assert.Contains(t, foundSchedules, schedules[0])
		assert.Contains(t, foundSchedules, schedules[1])
	})
}

func createJobSchedule(tx *sqlx.Tx, jobId uuid.UUID) repositories.JobSchedule {
	schedule := repositories.JobSchedule{
		Id:    uuid.New(),
		JobId: jobId,
		Cron:  "some-cron",
	}
	query := goqu.Insert(repositories.JobSchedulesTableName).Rows(&schedule)
	_ = repositories.Query(createContext(tx), query)

	return schedule
}
