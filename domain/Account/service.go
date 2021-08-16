package account

import (
	"github.com/Ammce/go-banking-core/dto/accountDTO"
	"github.com/Ammce/go-banking-core/errs"
)

type AccountService interface {
	CreateAccount(adto *accountDTO.CreateAccountDTO) (*accountDTO.CreateAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo AccountRepository
}

func (as DefaultAccountService) CreateAccount(caDTO *accountDTO.CreateAccountDTO) (*accountDTO.CreateAccountResponse, *errs.AppError) {
	validationError := caDTO.Validate()
	if validationError != nil {
		return nil, validationError
	}
	newAcc := NewAccount(caDTO.AccountType, caDTO.Amount, caDTO.CustomerID)
	acc, err := as.repo.Create(&newAcc)
	if err != nil {
		return nil, err
	}
	return acc.asResponseDto(), nil
}

func NewAccountService(repo AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
