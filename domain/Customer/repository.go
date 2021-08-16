package customer

import (
	"database/sql"

	"github.com/Ammce/go-banking-core/dto/customerDTO"
	"github.com/Ammce/go-banking-core/errs"
)

type CustomerRepository interface {
	Create(customer Customer) (*Customer, *errs.AppError)
	Update(id uint, customer Customer) (*Customer, *errs.AppError)
	FindAll(status string) ([]Customer, *errs.AppError)
	FindById(id uint) (*Customer, *errs.AppError)
	DeleteById(id uint) *errs.AppError
}

func (c Customer) AsResponseDto() *customerDTO.CustomerResponse {
	return &customerDTO.CustomerResponse{
		ID:          c.ID,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateOfBirth,
		Status:      c.Status,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
		DeletedAt:   sql.NullTime(c.DeletedAt),
	}
}

func NewCustomer(ccDTO customerDTO.CreateCustomer) Customer {
	return Customer{
		Name:        ccDTO.Name,
		City:        ccDTO.City,
		Zipcode:     ccDTO.Zipcode,
		DateOfBirth: ccDTO.DateofBirth,
		Status:      ccDTO.Status,
	}
}
