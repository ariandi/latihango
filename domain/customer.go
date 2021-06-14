package domain

import (
	"latihan1/dto"
	"latihan1/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) StatusText() string  {
	statusText := "active"
	if c.Status == "0" {
		statusText = "inactive"
	}

	return statusText
}

func (c Customer) ToDto() dto.CustomerResponse  {
	return dto.CustomerResponse{
		Id: c.Id,
		Name: c.Name,
		City: c.City,
		Zipcode: c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status: c.StatusText(),
	}
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}