package services

import (
	"errors"
	"project/internal/model"

	//"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func (s *Service) CompanyCreate(nc model.CreateCompany) (model.Company, error) {
	company := model.Company{CompanyName: nc.CompanyName, Adress: nc.Adress, Domain: nc.Domain}
	cu, err := s.r.CreateCompany(company)
	if err != nil {
		log.Error().Err(err).Msg("couldnot create user")
		return model.Company{}, errors.New("user creation failed")
	}

	return cu, nil
}

func (s *Service) GetAllCompanies() ([]model.Company, error) {

	AllCompanies, err := s.r.GetAllCompany()
	if err != nil {
		return nil, err
	}
	return AllCompanies, nil

}

func (s *Service) GetCompany(id int) (model.Company, error) {

	AllCompanies, err := s.r.GetCompany(id)
	if err != nil {
		return model.Company{}, err
	}
	return AllCompanies, nil

}

func (s *Service) JobCreate(nj model.CreateJob, id uint64) (model.Job, error) {
	job := model.Job{
		JobTitle:       nj.JobTitle,
		JobSalary:      nj.JobSalary,
		MinNP:          nj.MinNP,
		MaxNP:          nj.MaxNP,
		Budget:         nj.Budget,
		Description:    nj.Description,
		MinExp:         nj.MinExp,
		MaxMax:         nj.MaxMax,
		Qualification:  nj.Qualification,
		Shift:          nj.Shift,
		JobType:        nj.JobType,
		JobLocations:   nj.JobLocations,
		WorkMode:       nj.WorkMode,
		TechnologyStac: nj.TechnologyStac}
	cu, err := s.r.CreateJob(job)
	if err != nil {
		log.Error().Err(err).Msg("couldnot create user")
		return model.Job{}, errors.New("job creation failed")
	}

	return cu, nil
}

func (s *Service) GetJobs(id int) ([]model.Job, error) {
	AllCompanies, err := s.r.GetJobs(id)
	if err != nil {
		return nil, errors.New("job retreval failed")
	}
	return AllCompanies, nil
}

func (s *Service) GetAllJobs() ([]model.Job, error) {

	AllJobs, err := s.r.GetAllJobs()
	if err != nil {
		return nil, err
	}
	return AllJobs, nil

}
