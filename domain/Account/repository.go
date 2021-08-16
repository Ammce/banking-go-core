package account

import (
	"time"

	"github.com/Ammce/go-banking-core/errs"
)

type AccountRepository interface {
	Create(account *Account) (*Account, *errs.AppError)
}

func NewAccount(accountType string, amount float64, customerId uint) Account {
	currentTime := time.Now()
	return Account{
		OpeningDate: currentTime.Format("2006-01-02 15:04:05"),
		AccountType: accountType,
		Amount:      amount,
		CustomerID:  customerId,
		Status:      "active",
	}
}
