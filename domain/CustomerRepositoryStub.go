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
	customers := []Customer{
		{"1001", "Ammce", "Novi Pazar", "110011", "2000-01-01", "1"},
		{"1002", "Extra", "New York", "110011", "2000-01-01", "1"},
	}
	return CustomerRepositoryStub{customers}
}
