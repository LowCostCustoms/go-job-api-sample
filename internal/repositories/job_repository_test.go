package repositories_test

import (
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"go-scheduler-api/internal/context/vars"
	"go-scheduler-api/internal/repositories"
	"go-scheduler-api/internal/test/db"
	"testing"
)

func TestJobRepository_FindById(t *testing.T) {
	db.WithTestDatabase(func(tx *sqlx.Tx) {
		job := createJob(tx)
		repository := repositories.NewJobRepository()

		foundJob, err := repository.FindById(createContext(tx), job.Id)

		assert.NotNil(t, foundJob)
		assert.Nil(t, err)
		assert.Equal(t, *foundJob, job)
	})
}

func TestJobRepository_FindById_NotFound(t *testing.T) {
	db.WithTestDatabase(func(tx *sqlx.Tx) {
		repository := repositories.NewJobRepository()

		foundJob, err := repository.FindById(createContext(tx), uuid.New())

		assert.Nil(t, foundJob)
		assert.Nil(t, err)
	})
}

func TestJobRepository_FindAll(t *testing.T) {
	db.WithTestDatabase(func(tx *sqlx.Tx) {
		jobs := createManyJobs(tx, 3)
		repository := repositories.NewJobRepository()

		foundJobs, err := repository.FindAll(createContext(tx), repositories.JobQueryParams{})

		assert.Nil(t, err)
		assert.ElementsMatch(t, foundJobs, jobs)
	})
}

func TestJobRepository_FindAll_RespectsPagination(t *testing.T) {
	db.WithTestDatabase(func(tx *sqlx.Tx) {
		jobs := createManyJobs(tx, 3)
		repository := repositories.NewJobRepository()

		foundJobs, err := repository.FindAll(createContext(tx), repositories.JobQueryParams{
			Offset: 1,
			Limit:  2,
		})

		assert.Nil(t, err)
		assert.Len(t, foundJobs, 2)
		assert.Subset(t, jobs, foundJobs)
	})
}

func TestJobRepository_Count(t *testing.T) {
	db.WithTestDatabase(func(tx *sqlx.Tx) {
		jobCount := uint(3)
		createManyJobs(tx, jobCount)
		repository := repositories.NewJobRepository()

		count, err := repository.Count(createContext(tx), repositories.JobQueryParams{})

		assert.Nil(t, err)
		assert.Equal(t, count, jobCount)
	})
}

func TestJobRepository_Create(t *testing.T) {
	db.WithTestDatabase(func(tx *sqlx.Tx) {
		job := repositories.Job{
			Name: "some-job",
		}
		repository := repositories.NewJobRepository()
		ctx := createContext(tx)

		createdJob, err := repository.Create(ctx, job)

		assert.Nil(t, err)
		assert.NotNil(t, createdJob)
		assert.NotEqual(t, createdJob.Id, uuid.Nil)
		assert.Equal(t, createdJob.Name, job.Name)

		foundJob, err := repository.FindById(ctx, createdJob.Id)

		assert.Nil(t, err)
		assert.NotNil(t, foundJob)
		assert.Equal(t, createdJob, foundJob)
	})
}

func createContext(tx *sqlx.Tx) context.Context {
	return vars.WithQuerier(context.Background(), tx)
}

func createJob(tx *sqlx.Tx) repositories.Job {
	job := repositories.Job{
		Id:   uuid.New(),
		Name: fmt.Sprintf("job-%s", uuid.New()),
	}
	query := goqu.Insert(repositories.JobsTableName).Rows(&job)
	_ = repositories.Query(createContext(tx), query)

	return job
}

func createManyJobs(tx *sqlx.Tx, count uint) []repositories.Job {
	jobs := make([]repositories.Job, count)
	for idx := uint(0); idx != count; idx++ {
		jobs[idx] = createJob(tx)
	}

	return jobs
}
