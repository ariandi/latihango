package domain

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"latihan1/errs"
	"latihan1/logger"
	"strconv"
)

type TransactionRepositoryDb struct {
	client *sqlx.DB
}

func (d TransactionRepositoryDb) Save(t Transaction) (*Transaction, *errs.AppError)  {
	sqlInsert := "INSERT INTO transactions (account_id, transaction_type, amount, transaction_date) VALUES (?, ?, ?, ?)"
	result, err := d.client.Exec(sqlInsert, t.AccountId, t.TransactionType, t.Amount, t.TransactionDate)
	if err != nil {
		logger.Error("error while creating new trx : " + err.Error())
		return nil, errs.NewUnexpectedError("unexpect error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("error while getting last insert id for new trx : " + err.Error())
		return nil, errs.NewUnexpectedError("unexpect error from database")
	}

	t.TransactionId = strconv.FormatInt(id, 10)
	t.NewBalance = t.Amount
	return &t, nil
}

func (d TransactionRepositoryDb) FindById(accountId string) (*Account, *errs.AppError)  {
	AccountSql := "SELECT account_id, amount FROM accounts where account_id = ? and status = 1"
	var a Account
	err := d.client.Get(&a, AccountSql, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("account id not found")
			return nil, errs.NewNotFoundError("account id not found")
		} else {
			logger.Error("Error while scanning account " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return &a, nil
}

func NewTransactionRepositoryDb (dbClient *sqlx.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{dbClient}
}