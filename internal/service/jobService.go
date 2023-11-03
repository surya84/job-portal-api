package service

import (
	"context"
	"job-portal/internal/models"
)

func (r NewService) CreateJob(ctx context.Context, nj models.NewJob, cId int) (models.Job, error) {
	job, err := r.rs.CreateJ(ctx, nj, cId)
	return job, err
}

func (r NewService) ViewJob() ([]models.Job, error) {
	jobs, err := r.rs.ViewJobs()
	return jobs, err
}

func (r NewService) GetJobInfoByID(jId int) (models.Job, error) {
	job, err := r.rs.GetJobById(jId)
	return job, err
}

func (r NewService) ViewJobByCompanyId(cId int) ([]models.Job, error) {
	jobs, err := r.rs.ViewJobById(cId)
	return jobs, err
}
