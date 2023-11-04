package service

import (
	"context"
	"job-portal/internal/models"
)

func (r NewService) CreateJob(ctx context.Context, nj models.NewJob, cId int) (models.Job, error) {
	jobData, err := r.rs.CreateJ(ctx, nj, cId)
	if err != nil {
		return models.Job{}, err
	}
	return jobData, err
}

func (r NewService) ViewJob(ctx context.Context) ([]models.Job, error) {
	jobDetails, err := r.rs.ViewJobs(ctx)
	if err != nil {
		return []models.Job{}, err
	}
	return jobDetails, err
}

func (r NewService) GetJobInfoByID(jId int) (models.Job, error) {
	jobData, err := r.rs.GetJobById(jId)
	if err != nil {
		return models.Job{}, err
	}
	return jobData, err
}

func (r NewService) ViewJobByCompanyId(cId int) ([]models.Job, error) {
	jobData, err := r.rs.ViewJobById(cId)
	if err != nil {
		return []models.Job{}, err
	}
	return jobData, err
}
