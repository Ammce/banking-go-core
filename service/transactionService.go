package service

import (
	"github.com/Ammce/go-banking-core/domain"
	"github.com/Ammce/go-banking-core/errs"
)

type TransactionService interface {
	Create(t domain.Transaction) (*domain.Transaction, *errs.AppError)
}

type DefaultTransactionService struct {
	repo domain.TransactionRepositoryDB
}

func (ts DefaultTransactionService) Create(t domain.Transaction) (*domain.Transaction, *errs.AppError) {
	transaction, err := ts.repo.Create(t)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func NewTransactionService(repo domain.TransactionRepositoryDB) DefaultTransactionService {
	return DefaultTransactionService{
		repo: repo,
	}
}
