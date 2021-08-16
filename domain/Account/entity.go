package account

import (
	customer "github.com/Ammce/go-banking-core/domain/Customer"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
	CustomerID  string
	Customer    customer.Customer `gorm:"foreignKey:CustomerID"`
}
