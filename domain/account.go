package domain

import (
	"github.com/Dubjay18/gobank2/dto"
	"github.com/Dubjay18/gobank2/errs"
)

type Account struct {
	AccountId   string `db:"account_id"`
	CustomerId  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}

func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount < amount {
		return true
	}
	return false
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	ById(string) (*Account, *errs.AppError)
}
