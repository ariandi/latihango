package domain

import (
	"latihan1/dto"
	"latihan1/errs"
)

type Account struct {
	AccountId   string `db:"account_id" json:"account_id"`
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}

func (a Account) ToNewAccountRepositoryDto() dto.NewAccountResponse  {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}