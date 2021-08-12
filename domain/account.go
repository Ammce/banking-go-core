package domain

import (
	"github.com/Ammce/go-banking-core/dto"
	"github.com/Ammce/go-banking-core/errs"
	"github.com/google/uuid"
)

type Account struct {
	AccountId   uuid.UUID `json:"account_id"`
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      int64
	Status      string
}

type AccountRepository interface {
	Save(account *Account) (*Account, *errs.AppError)
}

func (a Account) ToAccountResponseDto() *dto.AccountResponse {
	return &dto.AccountResponse{AccountId: a.AccountId}
}

func NewAccount(customerId, accountType string, amount int64) Account {
	return Account{
		CustomerId:  customerId,
		OpeningDate: "2006-01-02 15:04:05",
		AccountType: accountType,
		Amount:      amount,
		Status:      "1",
	}
}
