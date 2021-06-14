package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"latihan1/errs"
	"latihan1/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError)  {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("customer not found")
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return &c, nil
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError)  {
	//var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)
	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}


	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	//customers := make([]Customer, 0)
	//err = sqlx.StructScan(rows, &customers)
	//if err != nil {
	//	logger.Error("Error while scanning customer " + err.Error())
	//	return nil, errs.NewUnexpectedError("unexpected database error")
	//}

	//for rows.Next() {
	//	var c Customer
	//	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	//	if err != nil {
	//		logger.Error("Error while scanning customer " + err.Error())
	//		return nil, errs.NewUnexpectedError("unexpected database error")
	//	}
	//
	//	customers = append(customers, c)
	//}

	return customers, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb  {
	return CustomerRepositoryDb{dbClient}
}
