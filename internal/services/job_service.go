package services

import (
	"context"
	"github.com/google/uuid"
	"go-scheduler-api/internal/api"
	"go-scheduler-api/internal/errors"
	"go-scheduler-api/internal/repositories"
)

type jobService struct {
	api.UnimplementedJobServiceServer

	jobRepository         repositories.JobRepository
	jobScheduleRepository repositories.JobScheduleRepository
	transactionManager    TransactionManager
}

func (s *jobService) ListJobs(ctx context.Context, request *api.ListJobsRequest) (*api.JobPageReply, error) {
	queryParams := repositories.JobQueryParams{
		Offset: uint(request.Offset),
		Limit:  uint(request.Limit),
	}
	count, err := s.jobRepository.Count(ctx, queryParams)
	if err != nil {
		return nil, err
	}

	jobs, err := s.jobRepository.FindAll(ctx, queryParams)
	if err != nil {
		return nil, err
	}

	jobReplies := make([]*api.JobReply, len(jobs))
	for idx, job := range jobs {
		jobReplies[idx] = &api.JobReply{
			Id:   job.Id.String(),
			Name: job.Name,
		}
	}

	return &api.JobPageReply{
		Count: uint32(count),
		Items: jobReplies,
	}, nil
}

func (s *jobService) GetJob(ctx context.Context, request *api.GetJobRequest) (*api.JobReply, error) {
	jobId, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, errors.NewBadRequestError("job id is invalid")
	}

	job, err := s.jobRepository.FindById(ctx, jobId)
	if err != nil {
		return nil, err
	}

	if job == nil {
		return nil, errors.NewNotFoundError("could not find a job")
	}

	schedules, err := s.jobScheduleRepository.FindByJobId(ctx, jobId)
	if err != nil {
		return nil, err
	}

	scheduleReplies := make([]*api.JobScheduleReply, len(schedules))
	for idx, schedule := range schedules {
		scheduleReplies[idx] = &api.JobScheduleReply{
			Id:    schedule.Id.String(),
			JobId: schedule.JobId.String(),
			Cron:  schedule.Cron,
		}
	}

	return &api.JobReply{
		Id:        job.Id.String(),
		Name:      job.Name,
		Schedules: scheduleReplies,
	}, nil
}

func (s *jobService) CreateJob(ctx context.Context, request *api.JobRequest) (*api.JobReply, error) {
	return WithTransaction(ctx, s.transactionManager, func(ctx context.Context) (*api.JobReply, error) {
		createdJob, err := s.jobRepository.Create(
			ctx,
			repositories.Job{
				Name: request.Name,
			},
		)
		if err != nil {
			return nil, err
		}

		jobSchedules := make([]repositories.JobSchedule, len(request.Schedules))
		for idx, schedule := range request.Schedules {
			jobSchedules[idx] = repositories.JobSchedule{
				JobId: createdJob.Id,
				Cron:  schedule.Cron,
			}
		}

		createdJobSchedules, err := s.jobScheduleRepository.CreateMany(ctx, jobSchedules...)
		if err != nil {
			return nil, err
		}

		jobScheduleReplies := make([]*api.JobScheduleReply, len(createdJobSchedules))
		for idx, schedule := range createdJobSchedules {
			jobScheduleReplies[idx] = &api.JobScheduleReply{
				Id:    schedule.Id.String(),
				JobId: schedule.JobId.String(),
				Cron:  schedule.Cron,
			}
		}

		return &api.JobReply{
			Id:        createdJob.Id.String(),
			Name:      createdJob.Name,
			Schedules: jobScheduleReplies,
		}, nil
	})
}

func NewJobServiceServer(
	jobRepository repositories.JobRepository,
	jobScheduleRepository repositories.JobScheduleRepository,
	transactionManager TransactionManager,
) api.JobServiceServer {
	return &jobService{
		jobRepository:         jobRepository,
		jobScheduleRepository: jobScheduleRepository,
		transactionManager:    transactionManager,
	}
}
