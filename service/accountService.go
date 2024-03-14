package service

import (
	"github.com/Dubjay18/gobank2/domain"
	"github.com/Dubjay18/gobank2/dto"
	"github.com/Dubjay18/gobank2/errs"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepositoryDB
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	response := newAccount.ToNewAccountResponseDto()
	return &response, nil
}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	if req.IsWithdrawal() {
		account, err := s.repo.ById(req.AccountId)

		if err != nil {
			return nil, err
		}
		if account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance")
		}
	}

	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	newTransaction, err := s.repo.SaveTransaction(t)
	if err != nil {
		return nil, err
	}
	response := newTransaction.ToDto()
	return &response, nil
}
func NewAccountService(repo domain.AccountRepositoryDB) DefaultAccountService {
	return DefaultAccountService{repo}
}
