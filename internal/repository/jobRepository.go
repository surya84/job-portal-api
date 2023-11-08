package repository

import (
	"context"
	"errors"
	"job-portal/internal/models"
)

func (s *Conn) CreateJ(ctx context.Context, nj models.NewJobRequest, cId int) (models.Job, error) {

	job := models.Job{
		Title:              nj.Title,
		Description:        nj.Description,
		Min_NoticePeriod:   nj.Min_NoticePeriod,
		Max_NoticePeriod:   nj.Max_NoticePeriod,
		Budget:             nj.Budget,
		Minimum_Experience: nj.Minimum_Experience,
		Maximum_Experience: nj.Maximum_Experience,
		CompanyId:          uint(cId),
	}

	//job.JobLocations = getLocations(nj.JobLocations)

	// for _, v := range nj.Qualifications {
	// 	tempData := models.Qualification{
	// 		Model: gorm.Model{
	// 			ID: uint(),
	// 		},
	// 	}
	// 	job.Qualifications = append(job.Qualifications, tempData)
	// }

	// for _, v := range nj.JobLocations {
	// 	tempData := models.JobLocation{
	// 		ID: int(v),
	// 	}
	// 	job.JobLocations = append(job.JobLocations, tempData)
	// }
	// for _, v := range nj.Shift {
	// 	tempData := models.Shift{
	// 		ID: int(v),
	// 	}
	// 	job.Shifts = append(job.Shifts, tempData)
	// }
	// for _, v := range nj.Technology_stack {
	// 	tempData := models.Technology{
	// 		ID: int(v),
	// 	}
	// 	job.Technology_stack = append(job.Technology_stack, tempData)
	// }
	// for _, v := range nj.Job_Type {
	// 	tempData := models.JobType{
	// 		ID: int(v),
	// 	}
	// 	job.JobTypes = append(job.JobTypes, tempData)
	// }
	// for _, v := range nj.WorkMode {
	// 	tempData := models.WorkMode{
	// 		ID: int(v),
	// 	}
	// 	job.WorkMode = append(job.WorkMode, tempData)
	// }

	tx := s.db.WithContext(ctx).Create(&job)

	if tx.Error != nil {
		return models.Job{}, errors.New("creation of job failed")
	}

	return job, nil
}

// func getLocations(locationIds []models.Application) (locationData []models.) {
// 	for _, v := range locationIds {
// 		tempLocation := models.Location{
// 			Model: gorm.Model{
// 				ID: v.PlaceId,
// 			},
// 		}
// 		locationData = append(locationData, tempLocation)
// 	}
// 	return locationData

// }
func (s *Conn) ViewJobs() ([]models.Job, error) {
	var jobs []models.Job

	err := s.db.Preload("JobLocations").Find(&jobs).Error

	if err != nil {
		return []models.Job{}, err
	}

	return jobs, nil
}
func (s *Conn) GetJobById(jId int) (models.Job, error) {
	var job models.Job
	tx := s.db.Where("ID = ?", jId)
	err := tx.Find(&job).Error
	if err != nil {
		return models.Job{}, errors.New("company not found")
	}
	return job, nil
}
func (s *Conn) ViewJobById(cId int) ([]models.Job, error) {
	var jobs []models.Job

	tx := s.db.Where("company_id =?", cId)
	err := tx.Find(&jobs).Error

	if err != nil {
		return []models.Job{}, errors.New("no jobs for that company")
	}

	return jobs, nil
}

func (s *Conn) GetJobProcessData(id int) (models.Job, error) {

	var jobData models.Job

	tx := s.db.Preload("Qualifications").Preload("Locations").Preload("Shifts").Preload("Technology_stack").Preload("WorkMode").Preload("JobTypes").Where("id", id)

	err := tx.Find(&jobData).Error

	if err != nil {
		return models.Job{}, err
	}

	return jobData, nil
}
