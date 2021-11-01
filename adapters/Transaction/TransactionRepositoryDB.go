package transaction

import (
	transactionDomain "github.com/Ammce/go-banking-core/domain/Transaction"
	"github.com/Ammce/go-banking-core/errs"
	"gorm.io/gorm"
)

type TransactionRepositoryDB struct {
	db *gorm.DB
}

func (tr TransactionRepositoryDB) Create(transaction transactionDomain.Transaction) (*transactionDomain.Transaction, *errs.AppError) {
	err := tr.db.Create(&transaction).Error
	if err != nil {
		return nil, errs.NewUnexpectedError("Error with creating Transaction")
	}
	return &transaction, nil
}

func NewTransactionRepositoryDB(db *gorm.DB) TransactionRepositoryDB {
	return TransactionRepositoryDB{
		db: db,
	}
}
