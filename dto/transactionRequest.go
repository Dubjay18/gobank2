package dto

import "github.com/Dubjay18/gobank2/errs"

const (
	Withdrawal = "withdrawal"
	Deposit    = "deposit"
)

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

func (t TransactionRequest) IsWithdrawal() bool {
	return t.TransactionType == Withdrawal
}

func (r TransactionRequest) Validate() *errs.AppError {
	if r.TransactionType != Withdrawal && r.TransactionType != Deposit {
		return errs.NewValidationError("Transaction type should be withdrawal or deposit")
	}
	if r.Amount < 0 {
		return errs.NewValidationError("Amount cannot be less than zero")
	}
	return nil
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	NewBalance      float64 `json:"new_balance"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}
