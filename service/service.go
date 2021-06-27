package service

import (
	"GoEvents/repository"
	"GoEvents/requests"
)

type Service interface {
	GetAllEmployees() ([]requests.AccountCreateRequest, error)
	GetAccount() (requests.AccountCreateRequest, error)
	CreateAccount() error
	UpdateAccount() error
	DeleteAccount() error
}

type service struct {
	repo repository.Repository
}

func NewRepository(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAllEmployees() ([]requests.AccountCreateRequest, error) {
	return s.repo.GetAllEmployees()
}

func (s *service) GetAccount() (requests.AccountCreateRequest, error) {
	panic("implement me")
}

func (s *service) CreateAccount() error {
	panic("implement me")
}

func (s *service) UpdateAccount() error {
	panic("implement me")
}

func (s *service) DeleteAccount() error {
	panic("implement me")
}
