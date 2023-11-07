package models

import (
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	ID                 uint   `gorm:"primaryKey;autoIncrement"`
	Title              string `json:"title"`
	CompanyID          uint   `json:"company_id"`
	Min_NoticePeriod   string `json:"min_np" validate:"required"`
	Max_NoticePeriod   string `json:"max_np" validate:"required"`
	Budget             string `json:"budget" validate:"required"`
	Description        string `json:"description" validate:"required"`
	Minimum_Experience string `json:"min_exp" validate:"required"`
	Maximum_Experience string `json:"max_exp" validate:"required"`

	Qualifications []Qualification `gorm:"many2many:job_qualifications;"`
	Shifts         []Shift         `gorm:"many2many:job_shifts;"`
	JobTypes       []JobType       `gorm:"many2many:job_job_types;"`

	JobLocations     []JobLocation `gorm:"many2many:job_location_jobs;"`
	Technology_stack []Technology  `gorm:"many2many:technology_jobs;"`
	WorkMode         []WorkMode    `gorm:"many2many:work_mode_jobs;"`
}

type Qualification struct {
	gorm.Model
	Name string
}

type Shift struct {
	gorm.Model
	Name string
}

type JobType struct {
	gorm.Model
	Name string
}

type JobLocation struct {
	gorm.Model
	Name string
}

type Technology struct {
	gorm.Model
	Name string
}

type WorkMode struct {
	gorm.Model
	Name string
}

type NewJob struct {
	Title              string          `json:"title" validate:"required"`
	CompanyID          uint            `json:"company_id"`
	Min_NoticePeriod   string          `json:"min_np" validate:"required"`
	Max_NoticePeriod   string          `json:"max_np" validate:"required"`
	Budget             string          `json:"budget" validate:"required"`
	JobLocations       []JobLocation   `json:"job_locations" validate:"required"`
	Technology_stack   []Technology    `json:"technology_stack" validate:"required"`
	WorkMode           []WorkMode      `json:"workmode" validate:"required"`
	Description        string          `json:"description" validate:"required"`
	Minimum_Experience string          `json:"min_exp" validate:"required"`
	Maximum_Experience string          `json:"max_exp" validate:"required"`
	Qualification      []Qualification `json:"qualification" validate:"required"`
	Shift              []Shift         `json:"shifts" validate:"required"`
	Job_Type           []JobType       `json:"job_type" validate:"required"`
}
