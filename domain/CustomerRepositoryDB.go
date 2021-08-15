package domain

import (
	"github.com/Ammce/go-banking-core/errs"
	"github.com/Ammce/go-banking-core/logger"
	"gorm.io/gorm"
)

type CustomerRepositoryDB struct {
	db *gorm.DB
}

func (cr CustomerRepositoryDB) Create(customer Customer) (*Customer, *errs.AppError) {
	err := cr.db.Create(&customer).Error
	if err != nil {
		return nil, errs.NewUnexpectedError("Creation of customer failed")
	}
	return &customer, nil
}

func (cr CustomerRepositoryDB) Update(id int32, customer Customer) (*Customer, *errs.AppError) {
	err := cr.db.Model(&Customer{}).Where("id = ?", id).Updates(Customer{Name: customer.Name, Zipcode: customer.Zipcode, City: customer.City}).Error
	if err != nil {
		return nil, errs.NewUnexpectedError("Update of customer failed")
	}
	return &customer, nil
}

func (cr CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := []Customer{}
	if status == "" {
		err = cr.db.Find(&customers).Error
	} else {
		err = cr.db.Find(&customers, "status = ?", status).Error
	}
	if err != nil {
		logger.Error("Error happened while getting the results")
		return nil, errs.NewNotFoundError("Customers not found")
	}
	return customers, nil
}

func (cr CustomerRepositoryDB) FindById(id int32) (*Customer, *errs.AppError) {
	var err error
	customer := Customer{}
	err = cr.db.Where("id = ?", id).First(&customer).Error
	if err != nil {
		logger.Error("Error happened while getting the results")
		return nil, errs.NewNotFoundError("Customer not found")
	}
	return &customer, nil
}

func (cr CustomerRepositoryDB) DeleteById(id int32) *errs.AppError {
	err := cr.db.Delete(&Customer{}, id).Error
	if err != nil {
		return errs.NewNotFoundError("Database deletition error")
	}
	return nil
}

func NewCustomerREpositoryDB(client *gorm.DB) CustomerRepositoryDB {

	return CustomerRepositoryDB{db: client}
}
