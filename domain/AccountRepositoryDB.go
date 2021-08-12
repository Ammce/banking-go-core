package domain

import (
	"github.com/Ammce/go-banking-core/errs"
	"github.com/Ammce/go-banking-core/logger"
	"gorm.io/gorm"
)

type AccountRepositoryDB struct {
	db *gorm.DB
}

func (ar AccountRepositoryDB) Save(a *Account) (*Account, *errs.AppError) {
	err := ar.db.Create(a).Error
	if err != nil {
		logger.Error("Error while creating account - " + err.Error())
		return nil, errs.NewUnexpectedError("Error while creating account")
	}
	return a, nil

}
