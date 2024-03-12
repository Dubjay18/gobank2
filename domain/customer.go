package domain

import "github.com/Dubjay18/gobank2/errs"

type Customer struct {
	Id          string `db:"customer_id"json:"id"`
	Name        string `json:"fullname"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `db:"date_of_birth"json:"date_of_birth"`
	Status      string `json:"status"`
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
