package domain

import (
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

// func (c Customer) statusToText() string {
// 	status := "active"
// 	if c.Status == "0" {
// 		status = "inactive"
// 	}
// 	return status
// }

// func (c Customer) AsDto() *dto.CustomerResponse {
// 	return &dto.CustomerResponse{
// 		Id:          c.Id,
// 		Name:        c.Name,
// 		City:        c.City,
// 		Zipcode:     c.Zipcode,
// 		DateofBirth: c.DateofBirth,
// 		Status:      c.statusToText(),
// 	}
// }

type CustomerRepository interface {
	Create(customer Customer) (*Customer, *errs.AppError)
	FindAll(status string) ([]Customer, *errs.AppError)
	FindById(id int32) (*Customer, *errs.AppError)
}
