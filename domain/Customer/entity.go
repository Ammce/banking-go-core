package customer

import (
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
