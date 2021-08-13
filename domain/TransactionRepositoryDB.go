package domain

import (
	"github.com/Ammce/go-banking-core/errs"
	"gorm.io/gorm"
)

type TransactionRepositoryDB struct {
	db *gorm.DB
}

func (r TransactionRepositoryDB) Create(t Transaction) (*Transaction, *errs.AppError) {
	err := r.db.Create(t).Error
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}
	return &t, nil
}
