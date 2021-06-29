package service

import (
	"GoEvents/repository"
	"GoEvents/requests"
)

type Service interface {
	GetAllEmployees() ([]requests.AccountCreateRequest, error)
	GetAccount(id string) (requests.AccountCreateRequest, error)
	CreateAccount(requests.AccountCreateRequest) (requests.AccountCreateRequest, error)
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

func (s *service) GetAccount(id string) (requests.AccountCreateRequest, error) {
	return s.repo.GetAccount(id)
}

func (s *service) CreateAccount(request requests.AccountCreateRequest) (requests.AccountCreateRequest, error) {
	return s.repo.CreateAccount(request)
}

func (s *service) UpdateAccount() error {
	panic("implement me")
}

func (s *service) DeleteAccount() error {
	panic("implement me")
}
