package domain

import (
	"fmt"

	"github.com/Ammce/go-banking-core/errs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomerRepositoryDB struct {
	db *gorm.DB
}

func (p CustomerRepositoryDB) FindAll() ([]Customer, *errs.AppError) {
	var err error
	customers := []Customer{}
	err = p.db.Find(&customers).Error
	if err != nil {
		fmt.Println("Error happened while getting the results")
		return nil, errs.NewNotFoundError("Customers not found")
	}
	return customers, nil
}

func (p CustomerRepositoryDB) FindById(id int32) (*Customer, *errs.AppError) {
	var err error
	customer := Customer{}
	err = p.db.Where("id = ?", id).First(&customer).Error
	if err != nil {
		fmt.Println("Error happened while getting the results " + err.Error())
		return nil, errs.NewNotFoundError("Customer not found")
	}
	return &customer, nil
}

func NewCustomerREpositoryDB() CustomerRepositoryDB {
	dsn := "host=localhost user=postgres password=postgres dbname=banking port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate()

	if err != nil {
		fmt.Println("error connectin go the database")
		panic(err)
	}

	return CustomerRepositoryDB{db}
}
