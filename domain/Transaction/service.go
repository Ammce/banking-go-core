package transaction

import (
	"fmt"

	"github.com/Ammce/go-banking-core/errs"
)

type TransactionService interface {
	CreateTransaction(t Transaction) (*Transaction, *errs.AppError)
}

type DefaultTransactionService struct {
	repo TransactionRepository
}

func (ts DefaultTransactionService) CreateTransaction(t Transaction) (*Transaction, *errs.AppError) {
	fmt.Println(t.AccountFromID)
	transaction, err := ts.repo.Create(t)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func NewTransactionService(repo TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{repo: repo}
}
