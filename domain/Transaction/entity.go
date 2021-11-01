package transaction

import (
	account "github.com/Ammce/go-banking-core/domain/Account"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Amount        float64
	AccountFrom   account.Account `gorm:"foreignKey:AccountFromID"`
	AccountFromID uint
	AccountTo     account.Account `gorm:"foreignKey:AccountToID"`
	AccountToID   uint
	Status        string
}
