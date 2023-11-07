package service

import (
	"context"
	"job-portal/internal/models"
	reflect "reflect"
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

func (r NewService) ProcessJob(ctx context.Context, id int, nj models.NewJob) (models.NewJob, error) {

	jobDetails, err := r.rp.GetJobProcessData(id)

	if err != nil {
		return models.NewJob{}, err
	}

	if areFieldsMatching(nj, jobDetails) {
		return models.NewJob{}, nil
	}

	// If fields do not match, return an error
	return models.NewJob{}, err
}

func areFieldsMatching(request models.NewJob, job models.Job) bool {
	return request.Title == job.Title &&
		request.CompanyID == job.CompanyID &&
		request.Min_NoticePeriod == job.Min_NoticePeriod &&
		request.Max_NoticePeriod == job.Max_NoticePeriod &&
		request.Budget == job.Budget &&
		request.Description == job.Description &&
		request.Minimum_Experience == job.Minimum_Experience &&
		request.Maximum_Experience == job.Maximum_Experience &&
		areSlicesMatching(request.JobLocations, job.JobLocations) &&
		areSlicesMatching(request.Qualification, job.Qualifications) &&
		areSlicesMatching(request.Job_Type, job.JobTypes) &&
		areSlicesMatching(request.Shift, job.Shifts) &&
		areSlicesMatching(request.WorkMode, job.WorkMode) &&
		areSlicesMatching(request.Technology_stack, job.Technology_stack)
}

func areSlicesMatching(request, db interface{}) bool {
	requestValue := reflect.ValueOf(request)
	dbValue := reflect.ValueOf(db)

	if requestValue.Kind() != dbValue.Kind() || requestValue.Len() != dbValue.Len() {
		return false
	}

	// Create a map for faster comparison
	dbMap := make(map[interface{}]bool)
	for i := 0; i < dbValue.Len(); i++ {
		dbMap[dbValue.Index(i).Interface()] = true
	}

	// Check if each item in the request exists in the db
	for i := 0; i < requestValue.Len(); i++ {
		if _, exists := dbMap[requestValue.Index(i).Interface()]; !exists {
			return false
		}
	}

	return true
}
