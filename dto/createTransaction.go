package dto

import (
	"github.com/Ammce/go-banking-core/errs"
	"github.com/google/uuid"
)

type CreateTransactionDTO struct {
	AccountID uuid.UUID `json:"account_id"`
	Amount    int64     `json:"amount"`
}

func (ct CreateTransactionDTO) Validate() *errs.AppError {
	if ct.Amount < 1 {
		return errs.NewValidationError("Amount of a transaction must be greater than 1")
	}
	return nil
}
