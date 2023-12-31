package repository

import (
	"errors"
	"project/internal/model"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) (Users, error) {
	if db == nil {
		return nil, errors.New("db connection not given")
	}

	return &Repo{db: db}, nil

}

//go:generate mockgen -source=userDao.go -destination=usergoDao_mock.go -package=repository

type Users interface {
	CreateUser(model.User) (model.User, error)
	FetchUserByEmail(string) (model.User, error)
	CreateCompany(model.Company) (model.Company, error)
	GetAllCompany() ([]model.Company, error)
	GetCompany(id int) (model.Company, error)
	CreateJob(j model.Job) (model.Job, error)
	GetJobs(id int) ([]model.Job, error)
	GetAllJobs() ([]model.Job, error)
}

func (r *Repo) CreateUser(u model.User) (model.User, error) {
	err := r.db.Create(&u).Error
	if err != nil {
		return model.User{}, err
	}
	return u, nil
}

func (r *Repo) FetchUserByEmail(s string) (model.User, error) {
	var u model.User
	tx := r.db.Where("email=?", s).First(&u)
	if tx.Error != nil {
		return model.User{}, nil
	}
	return u, nil

}
