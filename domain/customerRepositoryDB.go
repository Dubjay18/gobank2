package domain

import (
	"database/sql"
	"errors"
	"github.com/Dubjay18/gobank2/errs"
	"github.com/Dubjay18/gobank2/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	findAllCustomersQuery         = "SELECT * FROM customers"
	findAllCustomersByStatusQuery = "SELECT * FROM customers WHERE status = $1"
	findCustomerByIdQuery         = "SELECT * FROM customers WHERE customer_id = $1"
)

type CustomerRepositoryDB struct {
	db *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	var err error

	if status == "" {
		err = d.db.Select(&customers, findAllCustomersQuery)
	} else {
		err = d.db.Select(&customers, findAllCustomersByStatusQuery, status)

	}
	if err != nil {
		logger.Error("Error querying customers" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	var c Customer
	err := d.db.Get(&c, findCustomerByIdQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error scanning customers" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return &c, nil

}

func NewCustomerRepositoryDB(dbClient *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{dbClient}
}
