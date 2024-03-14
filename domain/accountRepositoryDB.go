package domain

import (
	"github.com/Dubjay18/gobank2/errs"
	"github.com/Dubjay18/gobank2/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type AccountRepositoryDB struct {
	db *sqlx.DB
}

func (d AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := `INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES ($1, $2, $3, $4, $5) RETURNING account_id`
	var lastId int64
	err := d.db.QueryRow(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status).Scan(&lastId)
	if err != nil {
		logger.Error("Error creating account: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}

	if lastId == 0 {
		logger.Error("Error getting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}
	a.AccountId = strconv.FormatInt(lastId, 10)

	return &a, nil
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{dbClient}
}
