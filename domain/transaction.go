package domain

import (
	"latihan1/dto"
	"latihan1/errs"
	"strings"
)

type Transaction struct {
	TransactionId   string
	AccountId       string
	TransactionType string
	Amount          float64
	TransactionDate string
	NewBalance      float64
}

type TransactionRepository interface {
	Save(Transaction) (*Transaction, *errs.AppError)
	FindById(accountId string) (*Account, *errs.AppError)
}

func (t Transaction) ToNewTransactionRepositoryDto() dto.NewTransactionResponse {
	return dto.NewTransactionResponse{
		TransactionId: t.TransactionId,
		AccountId: t.AccountId,
		NewBalance: t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}

func (t Transaction) CanWithDraw (amount float64) bool {
	if t.Amount > amount {
		return false
	}
	return true
}

func (t Transaction) IsWithDraw () bool {
	if strings.ToLower(t.TransactionType) == "withdraw" {
		return true
	}
	return false
}