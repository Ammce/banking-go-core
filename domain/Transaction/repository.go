package transaction

import "github.com/Ammce/go-banking-core/errs"

type TransactionRepository interface {
	Create(transaction Transaction) (*Transaction, *errs.AppError)
}
