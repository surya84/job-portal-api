package repository

import (
	"context"
	"errors"
	"job-portal/internal/models"

	"gorm.io/gorm"
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

	//job.Locations = getLocations(nj.Locations)

	for _, v := range nj.Qualifications {
		tempData := models.Qualification{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.Qualifications = append(job.Qualifications, tempData)
	}

	for _, v := range nj.Locations {
		tempData := models.Location{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.Locations = append(job.Locations, tempData)
	}
	for _, v := range nj.Shifts {
		tempData := models.Shift{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.Shifts = append(job.Shifts, tempData)
	}
	for _, v := range nj.Technologies {
		tempData := models.Technology{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.Technologies = append(job.Technologies, tempData)
	}
	// for _, v := range nj.JobTypes {
	// 	tempData := models.JobType{
	// 		Model: gorm.Model{
	// 			ID: v,
	// 		},
	// 	}
	// 	job.JobTypes = append(job.JobTypes, tempData)
	// }

	for _, v := range nj.JobTypes {
		tempData := models.JobType{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.JobTypes = append(job.JobTypes, tempData)
	}

	for _, v := range nj.WorkModes {
		tempData := models.WorkMode{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.WorkModes = append(job.WorkModes, tempData)
	}

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
