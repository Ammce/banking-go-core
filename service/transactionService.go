package service

import (
	"github.com/Ammce/go-banking-core/domain"
	"github.com/Ammce/go-banking-core/dto"
	"github.com/Ammce/go-banking-core/errs"
)

type TransactionService interface {
	Create(t dto.CreateTransactionDTO) (*domain.Transaction, *errs.AppError)
}

type DefaultTransactionService struct {
	repo domain.TransactionRepositoryDB
}

func (ts DefaultTransactionService) Create(t dto.CreateTransactionDTO) (*domain.Transaction, *errs.AppError) {
	validationError := t.Validate()
	if validationError != nil {
		return nil, validationError
	}
	// 2. Transform data to domain known value
	newTransaction := domain.NewTransaction(t.AccountID, t.Amount)
	createdTransaction, creationError := ts.repo.Create(newTransaction)
	if creationError != nil {
		return nil, creationError
	}
	return createdTransaction, nil
}

func NewTransactionService(repo domain.TransactionRepositoryDB) DefaultTransactionService {
	return DefaultTransactionService{
		repo: repo,
	}
}
