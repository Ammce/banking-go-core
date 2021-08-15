package dto

import "github.com/google/uuid"

type TransactionResponseDTO struct {
	TransactionID   uuid.UUID `json:"transaction_id"`
	AccountID       uuid.UUID `json:"account_id"`
	Amount          int64     `json:"amount"`
	TransactionDate string    `json:"transaction_date"`
}
