package dto

import (
	"latihan1/errs"
	"strings"
)

type NewTransactionRequest struct {
	AccountId       string  `json:"account_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}

func (r NewTransactionRequest) ValidateTrx() *errs.AppError {
	if r.Amount <= 0 {
		return errs.NewValidateError("amount cannot should be more than 0")
	}
	if strings.ToLower(r.TransactionType) != "withdraw" && strings.ToLower(r.TransactionType) != "deposit" {
		return errs.NewUnexpectedError("transaction type should be withdraw or deposit")
	}
	return nil
}
