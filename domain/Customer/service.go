package customer

import (
	"github.com/Ammce/go-banking-core/dto/customerDTO"
	"github.com/Ammce/go-banking-core/errs"
)

type CustomerService interface {
	CreateCustomer(customer customerDTO.CreateCustomer) (*customerDTO.CustomerResponse, *errs.AppError)
	UpdateCustomer(id uint, customer customerDTO.CreateCustomer) (*customerDTO.CustomerResponse, *errs.AppError)
	GetAllCustomers(status string) (*[]customerDTO.CustomerResponse, *errs.AppError)
	GetCustomerById(id uint) (*customerDTO.CustomerResponse, *errs.AppError)
	DeleteCustomerById(id uint) *errs.AppError
}

type DefaultCustomerService struct {
	repo CustomerRepository
}

func (cs DefaultCustomerService) CreateCustomer(customer customerDTO.CreateCustomer) (*customerDTO.CustomerResponse, *errs.AppError) {
	verr := customer.Validate()
	if verr != nil {
		return nil, verr
	}
	domainCustomer := NewCustomer(customer)
	c, err := cs.repo.Create(domainCustomer)
	if err != nil {
		return nil, err
	}
	return c.AsResponseDto(), nil
}

func (cs DefaultCustomerService) UpdateCustomer(id uint, customer customerDTO.CreateCustomer) (*customerDTO.CustomerResponse, *errs.AppError) {
	verr := customer.Validate()
	if verr != nil {
		return nil, verr
	}
	domainCustomer := NewCustomer(customer)
	c, err := cs.repo.Update(id, domainCustomer)
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

func (s DefaultCustomerService) GetCustomerById(id uint) (*customerDTO.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return c.AsResponseDto(), nil
}
func (s DefaultCustomerService) DeleteCustomerById(id uint) *errs.AppError {
	err := s.repo.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}

func NewCustomerService(repo CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
