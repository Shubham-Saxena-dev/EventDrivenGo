package service

import (
	"GoEvents/Publisher/repository"
	"GoEvents/Publisher/requests"
)

type Service interface {
	GetAllEmployees() ([]requests.AccountCreateRequest, error)
	GetAccount(id string) (requests.AccountCreateRequest, error)
	CreateAccount(requests.AccountCreateRequest) (requests.AccountCreateRequest, error)
	UpdateAccount(string, requests.AccountUpdateRequest) error
	DeleteAccount(string) error
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

func (s *service) UpdateAccount(id string, request requests.AccountUpdateRequest) error {
	return s.repo.UpdateAccount(id, request)
}

func (s *service) DeleteAccount(id string) error {
	return s.repo.DeleteAccount(id)
}
