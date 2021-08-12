package service

import (
	"github.com/Ammce/go-banking-core/domain"
	"github.com/Ammce/go-banking-core/dto"
	"github.com/Ammce/go-banking-core/errs"
)

type AccountService interface {
	CreateAccount(adto *dto.CreateAccountDTO) (*dto.AccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (as DefaultAccountService) CreateAccount(adto *dto.CreateAccountDTO) (*dto.AccountResponse, *errs.AppError) {
	newAcc := domain.Account{
		CustomerId:  adto.CustomerId,
		AccountType: adto.AccountType,
		Amount:      int64(adto.Amount),
		Status:      "1",
	}
	acc, err := as.repo.Save(&newAcc)
	if err != nil {
		return nil, err
	}
	return acc.ToAccountResponseDto(), nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
