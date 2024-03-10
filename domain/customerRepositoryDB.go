package domain

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type PostgresStorage struct {
	db *sql.DB
}

type CustomerRepositoryDB struct {
	db *sql.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {

	findPsql := "SELECT * FROM customers"
	rows, err := d.db.Query(findPsql)
	if err != nil {
		log.Println("Error querying customers", err)
		return nil, err
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error scanning customers", err)
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
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
