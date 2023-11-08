package models

import (
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Title              string  `json:"title"`
	Min_NoticePeriod   int     `json:"min_np" validate:"required"`
	Max_NoticePeriod   int     `json:"max_np" validate:"required"`
	Budget             float64 `json:"budget" validate:"required"`
	Description        string  `json:"description" validate:"required"`
	Minimum_Experience float64 `json:"min_exp" validate:"required"`
	Maximum_Experience float64 `json:"max_exp" validate:"required"`

	Qualifications []Qualification `gorm:"many2many:job_qualifications;"`
	Shifts         []Shift         `gorm:"many2many:job_shifts;"`
	JobTypes       []JobType       `gorm:"many2many:job_job_types;"`
	Locations      []Location      `gorm:"many2many:job_locations;"`
	Technologies   []Technology    `gorm:"many2many:job_technology;"`
	WorkModes      []WorkMode      `gorm:"many2many:job_work_modes;"`
	CompanyId      uint            `json:"companyId"`
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

type Location struct {
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

type NewJobRequest struct {
	Title              string  `json:"title"`
	Min_NoticePeriod   int     `json:"min_np" validate:"required"`
	Max_NoticePeriod   int     `json:"max_np" validate:"required"`
	Budget             float64 `json:"budget" validate:"required"`
	Description        string  `json:"description" validate:"required"`
	Minimum_Experience float64 `json:"min_exp" validate:"required"`
	Maximum_Experience float64 `json:"max_exp" validate:"required"`

	Qualifications []uint `gorm:"many2many:job_qualifications;"`
	Shifts         []uint `gorm:"many2many:job_shifts;"`
	JobTypes       []uint `gorm:"many2many:job_job_types;"`
	Locations      []uint `gorm:"many2many:job_locations;"`
	Technologies   []uint `gorm:"many2many:job_technology;"`
	WorkModes      []uint `gorm:"many2many:job_work_modes;"`
}

type NewJobResponse struct {
	ID uint
}

type ApplicationRequest struct {
	Title              string  `json:"title"`
	Min_NoticePeriod   int     `json:"min_np" validate:"required"`
	Max_NoticePeriod   int     `json:"max_np" validate:"required"`
	Budget             float64 `json:"budget" validate:"required"`
	Description        string  `json:"description" validate:"required"`
	Minimum_Experience float64 `json:"min_exp" validate:"required"`
	Maximum_Experience float64 `json:"max_exp" validate:"required"`

	Qualifications []uint `gorm:"many2many:job_qualifications;"`
	Shifts         []uint `gorm:"many2many:job_shifts;"`
	JobTypes       []uint `gorm:"many2many:job_job_types;"`
	Locations      []uint `gorm:"many2many:job_locations;"`
	Technologies   []uint `gorm:"many2many:job_technology;"`
	WorkModes      []uint `gorm:"many2many:job_work_modes;"`
}
