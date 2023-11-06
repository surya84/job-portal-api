package service

import (
	"context"
	"job-portal/internal/models"
)

func (r NewService) CreateJob(ctx context.Context, nj models.NewJob, cId int) (models.Job, error) {
	jobData, err := r.rp.CreateJ(ctx, nj, cId)
	if err != nil {
		return models.Job{}, err
	}
	return jobData, err
}

func (r NewService) ViewJob(ctx context.Context) ([]models.Job, error) {
	jobData, err := r.rp.ViewJobs()
	if err != nil {
		return []models.Job{}, err
	}
	return jobData, err
}

func (r NewService) GetJobInfoByID(ctx context.Context, jId int) (models.Job, error) {
	jobDetails, err := r.rp.GetJobById(jId)
	if err != nil {
		return models.Job{}, err
	}
	return jobDetails, err
}

func (r NewService) ViewJobByCompanyId(ctx context.Context, cId int) ([]models.Job, error) {
	jobDetails, err := r.rp.ViewJobById(cId)
	if err != nil {
		return []models.Job{}, err
	}
	return jobDetails, err
}
