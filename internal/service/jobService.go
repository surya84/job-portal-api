package service

import (
	"context"
	"fmt"
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

func (r NewService) ProcessJob(ctx context.Context, id int, nj []models.ApplicationRequest) (*[]models.ApplicationRequest, error) {

	jobDetails, err := r.rp.GetJobProcessData(id)

	// var newjob models.NewJob

	if err != nil {
		return &[]models.ApplicationRequest{}, err
	}
	jobs := []models.ApplicationRequest{}

	//wg := new(sync.WaitGroup)

	for _, val := range nj {
		if areFieldsMatching(&val, &jobDetails) {

			jobs = append(jobs, val)

		}
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
		request.Maximum_Experience == job.Maximum_Experience &&
		areQualificationsMatching(request.Qualifications, job.Qualifications) &&
		areLocationsEqual(request.Locations, job.Locations) &&
		areJobTypesEqual(request.JobTypes, job.JobTypes) &&
		areWorkModesEqual(request.WorkModes, job.WorkModes) &&
		areShiftsEqual(request.Locations, job.Shifts) &&
		areTechnologiesEqual(request.Technologies, job.Technologies)

}

func areLocationsEqual(requestLocations []uint, location []models.Location) bool {
	var q []uint
	var dataQ []uint
	q = append(q, requestLocations...)
	//fmt.Println(q)

	for _, val := range location {

		dataQ = append(dataQ, uint(val.ID))
	}
	//fmt.Println(dataQ)
	return areSlicesEqual(q, dataQ)
}

func areJobTypesEqual(requesTypes []uint, JobTypes []models.JobType) bool {

	var q []uint
	var dataQ []uint
	q = append(q, requesTypes...)
	fmt.Println(q)

	for _, val := range JobTypes {

		dataQ = append(dataQ, uint(val.ID))
	}
	return areSlicesEqual(q, dataQ)
}

func areShiftsEqual(requestShifts []uint, jobShifts []models.Shift) bool {

	var q []uint
	var dataQ []uint

	q = append(q, requestShifts...)

	for _, val := range jobShifts {

		dataQ = append(dataQ, uint(val.ID))
	}
	return areSlicesEqual(q, dataQ)
}

func areTechnologiesEqual(requesttech []uint, jobtech []models.Technology) bool {

	var q []uint
	var dataQ []uint

	q = append(q, requesttech...)

	for _, val := range jobtech {

		dataQ = append(dataQ, uint(val.ID))
	}
	return areSlicesEqual(q, dataQ)
}

func areWorkModesEqual(requestModes []uint, jobModes []models.WorkMode) bool {

	var q []uint
	var dataQ []uint

	q = append(q, requestModes...)

	for _, val := range jobModes {

		dataQ = append(dataQ, uint(val.ID))
	}
	return areSlicesEqual(q, dataQ)
}

func areQualificationsMatching(requestQualifications []uint, jobQualifications []models.Qualification) bool {

	var q []uint
	var dataQ []uint

	q = append(q, requestQualifications...)

	for _, val := range jobQualifications {

		dataQ = append(dataQ, uint(val.ID))
	}
	return areSlicesEqual(q, dataQ)
}

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

// func areQualificationsMatching(requestQualifications []uint, jobQualifications []models.Qualification) bool {
// 	// Convert the qualification IDs to a slice of uint
// 	requestQualificationIDs := make([]uint, len(requestQualifications))
// 	for i, v := range requestQualifications {
// 		requestQualificationIDs[i] = uint(v)
// 	}

// 	fmt.Println(requestQualificationIDs)

// 	jobQualificationIDs := make(map[uint]bool)
// 	for _, val := range jobQualifications {
// 		jobQualificationIDs[uint(val.ID)] = true
// 	}

// 	fmt.Println(jobQualificationIDs)
// 	for _, id := range requestQualificationIDs {
// 		if _, exists := jobQualificationIDs[id]; !exists {
// 			return false
// 		}
// 	}

// 	return true
// }
