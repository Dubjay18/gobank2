package domain

import (
	"database/sql"
	"errors"
	"github.com/Dubjay18/gobank2/errs"
	"github.com/Dubjay18/gobank2/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

const (
	insertAccountQry = "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES ($1, $2, $3, $4, $5) RETURNING account_id"
	findAccountQry   = "SELECT account_id, customer_id, opening_date, account_type, amount, status FROM accounts WHERE account_id = $1"
)

type AccountRepositoryDB struct {
	db *sqlx.DB
}

func (d AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := insertAccountQry
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

func (d AccountRepositoryDB) ById(id string) (*Account, *errs.AppError) {
	var a Account
	err := d.db.Get(&a, findAccountQry, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("account not found")
		} else {
			logger.Error("Error scanning customers" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &a, nil
}
func (d AccountRepositoryDB) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := d.db.Begin()
	if err != nil {
		logger.Error("Error starting a new transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}

	sqlInsert := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES ($1, $2, $3, $4) RETURNING transaction_id"
	var lastId int64
	err = tx.QueryRow(sqlInsert, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate).Scan(&lastId)

	if err != nil {
		logger.Error("Error creating account: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}

	if lastId == 0 {
		logger.Error("Error getting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}
	if t.IsWithdrawal() {
		_, err = tx.Exec("UPDATE accounts SET amount = amount - $1 WHERE account_id = $2", t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec("UPDATE accounts SET amount = amount + $1 WHERE account_id = $2", t.Amount, t.AccountId)
	}

	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, errs.NewUnexpectedError("unexpected error from database")
		}
		logger.Error("Error updating account balance: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}
	account, error := d.ById(t.AccountId)
	if error != nil {
		return nil, error
	}

	t.TransactionId = strconv.FormatInt(lastId, 10)
	t.Amount = account.Amount
	return &t, nil
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{dbClient}
}
