package domain

import "github.com/Ammce/go-banking-core/errs"

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      int64
	Status      string
}

type AccountRepository interface {
	Save(account *Account) (*Account, *errs.AppError)
}
