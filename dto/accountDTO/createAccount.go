package accountDTO

import "github.com/Ammce/go-banking-core/errs"

type CreateAccountDTO struct {
	CustomerID  uint    `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (ca CreateAccountDTO) Validate() *errs.AppError {
	if ca.Amount < 1000 {
		return errs.NewValidationError("Amount should be greater than 1000")
	}
	return nil
}
