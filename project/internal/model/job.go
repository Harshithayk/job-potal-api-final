package model

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	JobTitle       string           `json:"job_title" validate:"required"`
	JobSalary      string           `json:"job_salary" validate:"required"`
	Company        Company          `gorm:"ForeignKey:uid"`
	Uid            uint64           `JSON:"uid, omitempty"`
	MinNP          string           `json:"Min-NP" validate:"required"`
	MaxNP          string           `json:"Max-NP " validate:"required"`
	Budget         string           `json:"Budget" validate:"required"`
	Description    string           `json:"Description" validate:"required"`
	MinExp         string           `json:"MinExp" validate:"required"`
	MaxMax         string           `json:"MaxMax" validate:"required"`
	Qualification  []Qualification  `gorm:"many2many:Qualification;"`
	Shift          []Shift          `json:"Shift" gorm:"many2many:Shift;"`
	JobLocations   []JobLocations   `json:"JobLocations" gorm:"many2many:JobLocations;"`
	TechnologyStac []TechnologyStac `json:"TechnologyStac" gorm:"many2many:TechnologyStac;"`
	WorkMode       []WorkMode       `json:"WorkMode" gorm:"many2many:WorkMode;"`
	JobType        []JobType        `json:"JobType" gorm:"many2many:JobType;"`
}
type Qualification struct {
	gorm.Model
	Degree string `json:"Degree" validate:"required"`
}
type Shift struct {
	gorm.Model
	Shifttype string `json:"Shifttype" validate:"required"`
}
type JobType struct {
	gorm.Model
	Jobtype string `json:"Jobtype" validate:"required"`
}
type JobLocations struct {
	gorm.Model
	LocationName string `json:"LocationName" validate:"required"`
	Id           int    `json:"Id" validate:"required"`
}
type TechnologyStac struct {
	gorm.Model
	Skilltype string `json:"skill1" validate:"required"`
}
type WorkMode struct {
	gorm.Model
	WorkModetype string `json:"WorkMode" validate:"required"`
}
type Company struct {
	gorm.Model
	CompanyName string `json:"company_name" validate:"required"`
	Adress      string `json:"company_adress" validate:"required"`
	Domain      string `json:"domain" validate:"required"`
}

type CreateCompany struct {
	gorm.Model
	CompanyName string `json:"company_name" validate:"required"`
	Adress      string `json:"company_adress" validate:"required"`
	Domain      string `json:"domain" validate:"required"`
}

type CreateJob struct {
	JobTitle       string           `json:"job_title" validate:"required"`
	JobSalary      string           `json:"job_salary" validate:"required"`
	MinNP          string           `json:"Min-NP" validate:"required"`
	MaxNP          string           `json:"Max-NP " validate:"required"`
	Budget         string           `json:"Budget" validate:"required"`
	Description    string           `json:"Description" validate:"required"`
	MinExp         string           `json:"MinExp" validate:"required"`
	MaxMax         string           `json:"MaxMax" validate:"required"`
	Qualification  []Qualification  `json:"Qualification"`
	Shift          []Shift          `json:"Shift" validate:"required"`
	JobLocations   []JobLocations   `json:"JobLocations" validate:"required"`
	TechnologyStac []TechnologyStac `json:"TechnologyStac"  validate:"required"`
	WorkMode       []WorkMode       `json:"WorkMode" validate:"required"`
	JobType        []JobType        `json:"JobType"  validate:"required" `
}
