package services

import (
	"fmt"
	"latihan1/domain"
	"latihan1/dto"
	"latihan1/errs"
	"latihan1/logger"
	"time"
)

type TransactionService interface {
	NewTransaction(request dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

type DefaultTransactionService struct {
	repo domain.TransactionRepository
}

func (s DefaultTransactionService) NewTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)  {
	err := req.ValidateTrx()
	if err != nil {
		return nil, err
	}

	t := domain.Transaction{
		TransactionType: req.TransactionType,
		Amount: req.Amount,
		AccountId: req.AccountId,
	}

	if t.IsWithDraw() {
		account, err := s.repo.FindById(t.AccountId)
		if err != nil {
			return nil, err
		}
		logger.Info("enter withdraw")
		amountStr := fmt.Sprintf("%f", t.Amount)
		amountStr2 := fmt.Sprintf("%f", account.Amount) // s == "123.456000"
		logger.Info("t.amount : " + amountStr)
		logger.Info("account.amount : " + amountStr2)

		if !t.CanWithDraw(account.Amount) {
			return nil, errs.NewValidateError("Insufficient balance in the account")
		}
	}

	r := domain.Transaction{
		TransactionId: "",
		AccountId: req.AccountId,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
		TransactionType: req.TransactionType,
		Amount: req.Amount,
	}

	newTransaction, err := s.repo.Save(r)
	if err != nil {
		return nil, err
	}

	response := newTransaction.ToNewTransactionRepositoryDto()

	return &response, nil
}

func NewTransactionService(repo domain.TransactionRepository) DefaultTransactionService  {
	return DefaultTransactionService{repo}
}