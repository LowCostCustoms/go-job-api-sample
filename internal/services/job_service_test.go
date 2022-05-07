package services_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go-scheduler-api/internal/api"
	"go-scheduler-api/internal/errors"
	"go-scheduler-api/internal/repositories"
	"go-scheduler-api/internal/services"
	repositories_mocks "go-scheduler-api/internal/test/mocks/repositories"
	services_mocks "go-scheduler-api/internal/test/mocks/services"
	"testing"
)

func TestJobService_GetJob(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	job := &repositories.Job{
		Id:   uuid.New(),
		Name: "some-job-name",
	}
	schedules := []repositories.JobSchedule{
		{
			Id:    uuid.New(),
			JobId: job.Id,
			Cron:  "some-cron",
		},
		{
			Id:    uuid.New(),
			JobId: job.Id,
			Cron:  "other-cron",
		},
	}
	expected_reply := &api.JobReply{
		Id:   job.Id.String(),
		Name: job.Name,
		Schedules: []*api.JobScheduleReply{
			{
				Id:    schedules[0].Id.String(),
				JobId: schedules[0].JobId.String(),
				Cron:  schedules[0].Cron,
			},
			{
				Id:    schedules[1].Id.String(),
				JobId: schedules[1].JobId.String(),
				Cron:  schedules[1].Cron,
			},
		},
	}

	jobRepository := repositories_mocks.NewMockJobRepository(controller)
	jobRepository.EXPECT().FindById(gomock.Any(), job.Id).Times(1).Return(job, nil)

	jobScheduleRepository := repositories_mocks.NewMockJobScheduleRepository(controller)
	jobScheduleRepository.EXPECT().FindByJobId(gomock.Any(), job.Id).Times(1).Return(schedules, nil)

	jobService := services.NewJobServiceServer(jobRepository, jobScheduleRepository, nil)

	reply, err := jobService.GetJob(
		context.Background(),
		&api.GetJobRequest{
			Id: job.Id.String(),
		},
	)

	assert.Nil(t, err)
	assert.NotNil(t, reply)
	assert.Equal(t, reply, expected_reply)
}

func TestJobService_GetJob_InvalidJobId(t *testing.T) {
	jobService := services.NewJobServiceServer(nil, nil, nil)

	reply, err := jobService.GetJob(
		context.Background(),
		&api.GetJobRequest{
			Id: "some-id",
		},
	)

	assert.Nil(t, reply)
	assert.NotNil(t, err)
	assert.IsType(t, &errors.GenericError[errors.BadRequestTag]{}, err)
}

func TestJobService_GetJob_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	jobRepository := repositories_mocks.NewMockJobRepository(controller)
	jobRepository.EXPECT().FindById(gomock.Any(), gomock.Any()).Times(1).Return(nil, nil)

	jobService := services.NewJobServiceServer(jobRepository, nil, nil)

	reply, err := jobService.GetJob(
		context.Background(),
		&api.GetJobRequest{
			Id: uuid.New().String(),
		},
	)

	assert.Nil(t, reply)
	assert.NotNil(t, err)
	assert.IsType(t, &errors.GenericError[errors.NotFoundTag]{}, err)
}

func TestJobService_ListJobs(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	queryParams := repositories.JobQueryParams{
		Offset: 1,
		Limit:  2,
	}
	jobs := []repositories.Job{
		{
			Id:   uuid.New(),
			Name: "some-job-name",
		},
		{
			Id:   uuid.New(),
			Name: "other-job-name",
		},
	}
	count := uint(2)
	expectedReply := &api.JobPageReply{
		Count: uint32(count),
		Items: []*api.JobReply{
			{
				Id:   jobs[0].Id.String(),
				Name: jobs[0].Name,
			},
			{
				Id:   jobs[1].Id.String(),
				Name: jobs[1].Name,
			},
		},
	}

	jobRepository := repositories_mocks.NewMockJobRepository(controller)
	jobRepository.EXPECT().Count(gomock.Any(), queryParams).Times(1).Return(count, nil)
	jobRepository.EXPECT().FindAll(gomock.Any(), queryParams).Times(1).Return(jobs, nil)

	jobService := services.NewJobServiceServer(jobRepository, nil, nil)

	reply, err := jobService.ListJobs(
		context.Background(),
		&api.ListJobsRequest{
			Offset: uint32(queryParams.Offset),
			Limit:  uint32(queryParams.Limit),
		},
	)

	assert.Nil(t, err)
	assert.NotNil(t, reply)
	assert.Equal(t, reply, expectedReply)
}

func TestJobService_CreateJob(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	job := repositories.Job{
		Id:   uuid.New(),
		Name: "some-job-name",
	}
	jobSchedules := []repositories.JobSchedule{
		{
			Id:    uuid.New(),
			JobId: job.Id,
			Cron:  "some-cron",
		},
		{
			Id:    uuid.New(),
			JobId: job.Id,
			Cron:  "other-cron",
		},
	}
	expectedReply := &api.JobReply{
		Id:   job.Id.String(),
		Name: job.Name,
		Schedules: []*api.JobScheduleReply{
			{
				Id:    jobSchedules[0].Id.String(),
				JobId: jobSchedules[0].JobId.String(),
				Cron:  jobSchedules[0].Cron,
			},
			{
				Id:    jobSchedules[1].Id.String(),
				JobId: jobSchedules[1].JobId.String(),
				Cron:  jobSchedules[1].Cron,
			},
		},
	}

	jobRepository := repositories_mocks.NewMockJobRepository(controller)
	jobRepository.EXPECT().Create(
		gomock.Any(),
		repositories.Job{
			Name: job.Name,
		},
	).Times(1).Return(&job, nil)

	jobScheduleRepository := repositories_mocks.NewMockJobScheduleRepository(controller)
	jobScheduleRepository.EXPECT().CreateMany(
		gomock.Any(),
		repositories.JobSchedule{
			JobId: jobSchedules[0].JobId,
			Cron:  jobSchedules[0].Cron,
		},
		repositories.JobSchedule{
			JobId: jobSchedules[1].JobId,
			Cron:  jobSchedules[1].Cron,
		},
	).Times(1).Return(jobSchedules, nil)

	transactionManager := services_mocks.NewMockTransactionManager(controller)
	transactionManager.EXPECT().Begin(gomock.Any()).Times(1).Return(nil, services.NewNoopTransaction(), nil)

	jobService := services.NewJobServiceServer(jobRepository, jobScheduleRepository, transactionManager)

	reply, err := jobService.CreateJob(
		context.Background(),
		&api.JobRequest{
			Name: job.Name,
			Schedules: []*api.JobScheduleRequest{
				{
					Cron: jobSchedules[0].Cron,
				},
				{
					Cron: jobSchedules[1].Cron,
				},
			},
		},
	)

	assert.Nil(t, err)
	assert.NotNil(t, reply)
	assert.Equal(t, reply, expectedReply)
}
