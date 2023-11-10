package service

import (
	"context"
	"errors"
	"job-portal/internal/models"
	"sync"

	"github.com/rs/zerolog/log"
)

func (r NewService) CreateJob(ctx context.Context, nj models.NewJobRequest, cId int) (models.NewJobResponse, error) {
	jobData, err := r.rp.CreateJ(ctx, nj, cId)
	if err != nil {
		return models.NewJobResponse{}, err
	}
	jobId := models.NewJobResponse{
		ID: jobData.ID,
	}
	return jobId, err
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

func Compare(nj models.ApplicationRequest, job models.Job) (models.ApplicationRequest, error) {

	//fmt.Println("entering into compare function")
	var count int
	err := errors.New("")
	if nj.Budget > job.Budget {

		return models.ApplicationRequest{}, err
	}

	if nj.NoticePeriod >= job.Min_NoticePeriod && nj.NoticePeriod <= job.Max_NoticePeriod {

		//log.Info().Str("Min_NP", "true").Send()
		count++
	} else {
		log.Info().Str("Min_NP", "false").Send()
		return models.ApplicationRequest{}, err
	}

	if nj.Experience < job.Minimum_Experience {

		return models.ApplicationRequest{}, err

	}
	//comparing job criteria locations and application criteria locations
	var loc_job []uint
	var loc_app []uint
	for _, v := range job.Locations {
		loc_job = append(loc_job, v.ID)
	}
	loc_app = nj.Locations
	if sliceContainsAtLeastOne(loc_job, loc_app) {
		count++
	}

	//comparing job criteria technologystack and application criteria technologystack
	var tech_job []uint
	var tech_app []uint
	for _, v := range job.Technologies {
		tech_job = append(tech_job, v.ID)
	}
	tech_app = nj.Technologies
	if sliceContainsAtLeastOne(tech_job, tech_app) {
		count++
	}

	//comparing job criteria technologystack and application criteria technologystack
	var mode_job []uint
	var mode_app []uint
	for _, v := range job.WorkModes {
		mode_job = append(mode_job, v.ID)
	}
	mode_app = nj.WorkModes
	if sliceContainsAtLeastOne(mode_job, mode_app) {
		count++
	}

	//comparing job criteria qualification and application criteria qualification
	var q_job []uint
	var q_app []uint
	for _, v := range job.Qualifications {
		q_job = append(q_job, v.ID)
	}
	q_app = nj.Qualifications
	if sliceContainsAtLeastOne(q_job, q_app) {
		count++
	}

	//comparing job criteria shifts and application criteria shifts
	var shift_job []uint
	var shift_app []uint
	for _, v := range job.Shifts {
		shift_job = append(shift_job, v.ID)
	}
	shift_app = nj.Shifts
	if sliceContainsAtLeastOne(shift_job, shift_app) {

		count++
	}

	//comparing job criteria technologystack and application criteria technologystack
	var type_job []uint
	var type_app []uint
	for _, v := range job.JobTypes {
		type_job = append(type_job, v.ID)
	}
	type_app = nj.JobTypes
	if sliceContainsAtLeastOne(type_job, type_app) {
		count++
	}

	if count >= 4 {
		return nj, nil
	}

	return models.ApplicationRequest{}, err
}

// function to check the slices
func sliceContainsAtLeastOne(slice, subSlice []uint) bool {
	for _, v := range subSlice {
		for _, s := range slice {
			if v == s {
				return true
			}
		}
	}
	return false

}

func (r NewService) ProcessJob(ctx context.Context, nj []models.ApplicationRequest) ([]models.ApplicationRequest, error) {

	var wg sync.WaitGroup
	userChannel := make(chan models.ApplicationRequest, len(nj))
	for _, application := range nj {
		wg.Add(1)
		go func(application models.ApplicationRequest) {
			defer wg.Done()
			jid := application.Id

			job, err := r.rp.GetJobProcessData(jid)

			if err != nil {
				return
			}

			user, err := Compare(application, job)

			if err != nil {
				return
			}

			// response, err := compare(nj, job)
			// if err != nil {
			// 	return []models.ApplicationRequest{}, err
			// }

			// result = append(result, response)

			userChannel <- user
		}(application)

	}

	go func() {
		wg.Wait()
		close(userChannel)
	}()

	var users []models.ApplicationRequest
	for user := range userChannel {
		users = append(users, user)
	}

	return users, nil
}
