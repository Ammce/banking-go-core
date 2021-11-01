package transaction

import (
	account "github.com/Ammce/go-banking-core/domain/Account"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Amount        float64
	AccountFrom   account.Account `gorm:"foreignKey:AccountFromID"`
	AccountFromID uint            `json:"account_from_id"`
	AccountTo     account.Account `gorm:"foreignKey:AccountToID"`
	AccountToID   uint            `json:"account_to_id"`
	Status        string
}
