package domain

import (
	"fmt"

	"github.com/Ammce/go-banking-core/errs"
	"github.com/Ammce/go-banking-core/logger"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountRepositoryDB struct {
	db *gorm.DB
}

func (ar AccountRepositoryDB) Save(a *Account) (*Account, *errs.AppError) {
	ac := Account{
		AccountId:   uuid.New(),
		CustomerId:  a.CustomerId,
		OpeningDate: a.OpeningDate,
		AccountType: a.AccountType,
		Amount:      a.Amount,
		Status:      a.Status,
	}
	fmt.Printf("%+v\n", ac)
	result := ar.db.Create(&ac)
	if result.Error != nil {
		logger.Error("Error while creating account - " + result.Error.Error())
		return nil, errs.NewUnexpectedError("Error while creating account")
	}

	fmt.Printf("%+v\n", ac)
	return &ac, nil

}

func NewAccountingRepositoryDB(client *gorm.DB) AccountRepositoryDB {
	return AccountRepositoryDB{db: client}
}
