package repository

import (
	"context"
	"errors"
	"job-portal/internal/models"

	"gorm.io/gorm"
)

func (s *Conn) CreateJ(ctx context.Context, newJob models.NewJobRequest, cId int) (models.NewJobResponse, error) {

	job := models.Job{
		Title:              newJob.Title,
		Description:        newJob.Description,
		Min_NoticePeriod:   *newJob.Min_NoticePeriod,
		Max_NoticePeriod:   *newJob.Max_NoticePeriod,
		Budget:             newJob.Budget,
		Minimum_Experience: newJob.Minimum_Experience,
		Maximum_Experience: newJob.Maximum_Experience,
		CompanyId:          uint(cId),
	}

	//job.Locations = getLocations(newJob.Locations)

	for _, v := range newJob.Qualifications {
		tempData := models.Qualification{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.Qualifications = append(job.Qualifications, tempData)
	}

	for _, v := range newJob.Locations {
		tempData := models.Location{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.Locations = append(job.Locations, tempData)
	}
	for _, v := range newJob.Shifts {
		tempData := models.Shift{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.Shifts = append(job.Shifts, tempData)
	}
	for _, v := range newJob.Technologies {
		tempData := models.Technology{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.Technologies = append(job.Technologies, tempData)
	}

	for _, v := range newJob.JobTypes {
		tempData := models.JobType{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.JobTypes = append(job.JobTypes, tempData)
	}

	for _, v := range newJob.WorkModes {
		tempData := models.WorkMode{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.WorkModes = append(job.WorkModes, tempData)
	}

	tx := s.db.WithContext(ctx).Create(&job)

	if tx.Error != nil {
		return models.NewJobResponse{}, errors.New("creation of job failed")
	}

	jobId := models.NewJobResponse{
		ID: job.ID,
	}

	return jobId, nil
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

	err := s.db.Preload("Locations").Preload("Qualifications").Preload("Locations").Preload("Shifts").Preload("Technologies").Preload("WorkModes").Preload("JobTypes").Find(&jobs).Error

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

	tx := s.db.Preload("Qualifications").Preload("Locations").Preload("Shifts").Preload("Technologies").Preload("WorkModes").Preload("JobTypes").Where("id", id)

	err := tx.Find(&jobData).Error

	if err != nil {
		return models.Job{}, err
	}

	return jobData, nil
}
