package domain

import (
	"database/sql"

	"github.com/Ammce/go-banking-core/dto/customerDTO"
	"github.com/Ammce/go-banking-core/errs"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
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

type CustomerRepository interface {
	Create(customer Customer) (*Customer, *errs.AppError)
	Update(id int32, customer Customer) (*Customer, *errs.AppError)
	FindAll(status string) ([]Customer, *errs.AppError)
	FindById(id int32) (*Customer, *errs.AppError)
	DeleteById(id int32) *errs.AppError
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
