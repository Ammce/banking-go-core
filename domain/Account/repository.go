package account

import (
	"database/sql"
	"time"

	"github.com/Ammce/go-banking-core/dto/accountDTO"
	"github.com/Ammce/go-banking-core/errs"
)

type AccountRepository interface {
	Create(account *Account) (*Account, *errs.AppError)
}

func (a Account) asResponseDto() *accountDTO.CreateAccountResponse {
	return &accountDTO.CreateAccountResponse{
		ID:          a.ID,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
		DeletedAt:   sql.NullTime(a.DeletedAt),
		OpeningDate: a.OpeningDate,
		AccountType: a.AccountType,
		Amount:      a.Amount,
		Status:      a.Status,
		CustomerID:  a.CustomerID,
	}
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
