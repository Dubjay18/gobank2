package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) findAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "John", "New York", "10001", "2000-01-01", "active"},
		{"1002", "John", "New York", "10001", "2000-01-01", "active"},
		{"1003", "John", "New York", "10001", "2000-01-01", "active"},
	}
	return CustomerRepositoryStub{customers}
}
