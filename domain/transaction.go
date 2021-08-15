package domain

import (
	"time"

	"github.com/Ammce/go-banking-core/errs"
	"github.com/google/uuid"
)

type Transaction struct {
	TransactionID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"transaction_id"`
	AccountID       uuid.UUID `json:"account_id"`
	Amount          int64     `json:"amount"`
	TransactionDate string    `json:"transaction_date"`
}

type TransactionRepository interface {
	Create(t Transaction) (*Transaction, *errs.AppError)
}

func NewTransaction(AccountID uuid.UUID, Amount int64) Transaction {
	currentTime := time.Now()
	return Transaction{
		AccountID:       AccountID,
		Amount:          Amount,
		TransactionDate: currentTime.Format("2006-01-02 15:04:05"),
	}
}
