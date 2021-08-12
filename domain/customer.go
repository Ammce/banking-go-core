package domain

import "github.com/Ammce/go-banking-core/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	FindById(id int32) (*Customer, *errs.AppError)
}
