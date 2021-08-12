package domain

import (
	"github.com/Ammce/go-banking-core/errs"
	"github.com/Ammce/go-banking-core/logger"
	"gorm.io/gorm"
)

type CustomerRepositoryDB struct {
	db *gorm.DB
}

func (p CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := []Customer{}
	if status == "" {
		err = p.db.Find(&customers).Error
	} else {
		err = p.db.Find(&customers, "status = ?", status).Error
	}
	if err != nil {
		logger.Error("Error happened while getting the results")
		return nil, errs.NewNotFoundError("Customers not found")
	}
	return customers, nil
}

func (p CustomerRepositoryDB) FindById(id int32) (*Customer, *errs.AppError) {
	var err error
	customer := Customer{}
	err = p.db.Where("id = ?", id).First(&customer).Error
	if err != nil {
		logger.Error("Error happened while getting the results")
		return nil, errs.NewNotFoundError("Customer not found")
	}
	return &customer, nil
}

func NewCustomerREpositoryDB(client *gorm.DB) CustomerRepositoryDB {

	return CustomerRepositoryDB{db: client}
}
