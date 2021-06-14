package services

import (
	"latihan1/domain"
	"latihan1/dto"
	"latihan1/errs"
	"time"
)

type AccountService interface {
	NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)  {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.Account{
		AccountId: "",
		CustomerId: req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount: req.Amount,
		Status: "1",
	}

	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountRepositoryDto()

	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService  {
	return DefaultAccountService{repo}
}