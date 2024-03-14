package domain

import (
	"github.com/jmoiron/sqlx"
)

type TransactionRepositoryDB struct {
	db *sqlx.DB
}

func NewTransactionRepositoryDB(dbClient *sqlx.DB) TransactionRepositoryDB {
	return TransactionRepositoryDB{dbClient}
}
