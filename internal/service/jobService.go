package service

import (
	"context"
	"job-portal/internal/models"
)

func (r NewService) CreateJob(ctx context.Context, nj models.NewJobRequest, cId int) (models.Job, error) {
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

func (r NewService) ProcessJob(ctx context.Context, id int, nj models.ApplicationRequest) (*[]models.ApplicationRequest, error) {

	jobDetails, err := r.rp.GetJobProcessData(id)

	// var newjob models.NewJob

	if err != nil {
		return &[]models.ApplicationRequest{}, err
	}
	jobs := []models.ApplicationRequest{}

	if areFieldsMatching(&nj, &jobDetails) {

		jobs = append(jobs, nj)

	}

	// If fields do not match, return an error
	return &jobs, err
}

func areFieldsMatching(request *models.ApplicationRequest, job *models.Job) bool {
	return request.Title == job.Title &&
		//request.companyId == job.CompanyId &&
		request.Min_NoticePeriod == job.Min_NoticePeriod &&
		request.Max_NoticePeriod == job.Max_NoticePeriod &&
		request.Budget == job.Budget &&
		request.Description == job.Description &&
		request.Minimum_Experience == job.Minimum_Experience &&
		request.Maximum_Experience == job.Maximum_Experience
	//areQualificationsMatching(request, job)
	//check(request.Qualification, job.Qualifications)
	//check2()

}

// func areQualificationsMatching(request *models.Application, jobs *models.Job) bool {

// 	var q []uint
// 	var dataQ []uint

// 	for _, v := range request.Qualification {
// 		q = append(q, v)
// 	}
// 	fmt.Println(q)

// 	for _, val := range jobs.Qualifications {

// 		fmt.Print(val.ID)

// 		dataQ = append(dataQ, uint(val.ID))
// 	}
// 	fmt.Println(dataQ)

// 	//equal := reflect.DeepEqual(q, dataQ)

// 	if areSlicesEqual(q, dataQ) {
// 		return true
// 	}

// 	return false

// }

func areSlicesEqual(slice1, slice2 []uint) bool {

	if len(slice1) != len(slice2) {
		return false
	}

	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
