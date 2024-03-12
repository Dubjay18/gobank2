package domain

import (
	"database/sql"
	"errors"
	"github.com/Dubjay18/gobank2/errs"
	"github.com/Dubjay18/gobank2/logger"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDB struct {
	db *sql.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {

	findPsql := "SELECT * FROM customers"
	rows, err := d.db.Query(findPsql)
	if err != nil {
		logger.Error("Error querying customers" + err.Error())
		return nil, err
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.DateOfBirth, &c.City, &c.Zipcode, &c.Status)
		if err != nil {
			logger.Error("Error scanning customers" + err.Error())
			return nil, err
		}

		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	findPsql := "SELECT * FROM customers WHERE customer_id = $1"
	row := d.db.QueryRow(findPsql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.DateOfBirth, &c.City, &c.Zipcode, &c.Status)
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

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	constr := "user=postgres dbname=goBank2 password=qwertyuiop sslmode=disable"
	db, err := sql.Open("postgres", constr)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)

	}
	return CustomerRepositoryDB{db}
}
