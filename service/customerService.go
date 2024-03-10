package service

import "github.com/Dubjay18/gobank2/domain"

// CustomerService is an interface that defines the methods that the service layer will use to interact with the domain layer
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repo domain.CustomerRepositoryDB) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
