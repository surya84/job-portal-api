package models

import (
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	//ID                 uint   `gorm:"primaryKey;autoIncrement"`
	Title              string `json:"title"`
	CompanyID          uint   `json:"company_id"`
	Min_NoticePeriod   int    `json:"min_np" validate:"required"`
	Max_NoticePeriod   int    `json:"max_np" validate:"required"`
	Budget             int    `json:"budget" validate:"required"`
	Description        string `json:"description" validate:"required"`
	Minimum_Experience int    `json:"min_exp" validate:"required"`
	Maximum_Experience int    `json:"max_exp" validate:"required"`

	Qualifications []Qualification `gorm:"many2many:job_qualifications;"`
	Shifts         []Shift         `gorm:"many2many:job_shifts;"`
	JobTypes       []JobType       `gorm:"many2many:job_job_types;"`

	JobLocations     []JobLocation `gorm:"many2many:job_location_jobs;"`
	Technology_stack []Technology  `gorm:"many2many:technology_jobs;"`
	WorkMode         []WorkMode    `gorm:"many2many:work_mode_jobs;"`
}

type Qualification struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Shift struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type JobType struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type JobLocation struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Technology struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type WorkMode struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type NewJob struct {
	Title              string          `json:"title" validate:"required"`
	CompanyID          uint            `json:"company_id"`
	Min_NoticePeriod   int             `json:"min_np" validate:"required"`
	Max_NoticePeriod   int             `json:"max_np" validate:"required"`
	Budget             int             `json:"budget" validate:"required"`
	JobLocations       []JobLocation   `json:"job_locations" validate:"required"`
	Technology_stack   []Technology    `json:"technology_stack" validate:"required"`
	WorkMode           []WorkMode      `json:"workmode" validate:"required"`
	Description        string          `json:"description" validate:"required"`
	Minimum_Experience int             `json:"min_exp" validate:"required"`
	Maximum_Experience int             `json:"max_exp" validate:"required"`
	Qualification      []Qualification `json:"qualification" validate:"required"`
	Shift              []Shift         `json:"shifts" validate:"required"`
	Job_Type           []JobType       `json:"job_type" validate:"required"`
}
