package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func (s CustomerRepositoryStub) FindById(id int32) (*Customer, error) {
	return &s.customers[0], nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{}
	return CustomerRepositoryStub{customers}
}
