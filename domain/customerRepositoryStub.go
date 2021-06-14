package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error)  {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer {
		{"1", "Ari", "Bogor", "20924", "1987-14-11", "1"},
		{"2", "Fio", "Bogor", "20923", "2014-05-01", "1"},
	}

	return CustomerRepositoryStub{customers}
}