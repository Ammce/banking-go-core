package domain

import (
	"github.com/Ammce/go-banking-core/errs"
	"github.com/google/uuid"
)

type Transaction struct {
	TransactionID   uuid.UUID `json:"transaction_id"`
	AccountID       uuid.UUID ` json:"account_id"`
	Amount          int64     `json:"amount"`
	TransactionDate string    `json:"transaction_date"`
}

type TransactionRepository interface {
	Create(t Transaction) (*Transaction, *errs.AppError)
}