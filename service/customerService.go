package service

import (
	"github.com/Ammce/go-banking-core/domain"
	"github.com/Ammce/go-banking-core/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]domain.Customer, *errs.AppError)
	GetCustomerById(id int32) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	customersResponse := make([]domain.Customer, 0)
	for _, customer := range customers {
		customersResponse = append(customersResponse, customer)
	}
	return customersResponse, nil
}

func (s DefaultCustomerService) GetCustomerById(id int32) (*domain.Customer, *errs.AppError) {
	c, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
