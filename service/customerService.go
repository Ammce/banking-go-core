package service

import (
	"github.com/Ammce/go-banking-core/domain"
	"github.com/Ammce/go-banking-core/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomerById(id int32) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomerById(id int32) (*domain.Customer, *errs.AppError) {
	return s.repo.FindById(id)
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
