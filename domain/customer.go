package domain

import "github.com/Dubjay18/gobank2/errs"

type Customer struct {
	Id          string `json:"id"`
	Name        string `json:"fullname"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
