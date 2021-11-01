package transaction

import (
	account "github.com/Ammce/go-banking-core/domain/Account"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	CreatedAt  string
	Amount     float64
	AcountFrom account.Account `gorm:"foreignKey:AccountID"`
	AcountTo   account.Account `gorm:"foreignKey:AccountID"`
	Status     string
}
