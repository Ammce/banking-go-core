package accountAdapter

import (
	account "github.com/Ammce/go-banking-core/domain/Account"
	"github.com/Ammce/go-banking-core/errs"
	"github.com/Ammce/go-banking-core/logger"
	"gorm.io/gorm"
)

type AccountRepositoryDB struct {
	db *gorm.DB
}

func (ar AccountRepositoryDB) Create(a *account.Account) (*account.Account, *errs.AppError) {
	result := ar.db.Create(&a)
	if result.Error != nil {
		logger.Error("Error while creating account - " + result.Error.Error())
		return nil, errs.NewUnexpectedError("Error while creating account")
	}

	return a, nil

}

func NewAccountRepositoryDB(client *gorm.DB) AccountRepositoryDB {
	return AccountRepositoryDB{db: client}
}
