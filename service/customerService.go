package service

import (
	"github.com/Ammce/go-banking-core/domain"
	"github.com/Ammce/go-banking-core/dto/customerDTO"
	"github.com/Ammce/go-banking-core/errs"
)

type CustomerService interface {
	CreateCustomer(customer customerDTO.CreateCustomer) (*customerDTO.CustomerResponse, *errs.AppError)
	GetAllCustomers(status string) (*[]customerDTO.CustomerResponse, *errs.AppError)
	GetCustomerById(id int32) (*customerDTO.CustomerResponse, *errs.AppError)
	DeleteCustomerById(id int32) *errs.AppError
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (cs DefaultCustomerService) CreateCustomer(customer customerDTO.CreateCustomer) (*customerDTO.CustomerResponse, *errs.AppError) {
	verr := customer.Validate()
	if verr != nil {
		return nil, verr
	}
	domainCustomer := domain.NewCustomer(customer)
	c, err := cs.repo.Create(domainCustomer)
	if err != nil {
		return nil, err
	}
	return c.AsResponseDto(), nil
}

func (s DefaultCustomerService) GetAllCustomers(status string) (*[]customerDTO.CustomerResponse, *errs.AppError) {
	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	customersResponse := make([]customerDTO.CustomerResponse, 0)
	for _, customer := range customers {
		customersResponse = append(customersResponse, *customer.AsResponseDto())
	}
	return &customersResponse, nil
}

func (s DefaultCustomerService) GetCustomerById(id int32) (*customerDTO.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return c.AsResponseDto(), nil
}
func (s DefaultCustomerService) DeleteCustomerById(id int32) *errs.AppError {
	err := s.repo.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
