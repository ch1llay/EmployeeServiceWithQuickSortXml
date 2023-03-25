package repository

import "EmployeeServiceWithQuickSortXml/Models"

type EmployeeRepository struct {
	connectionString string
}

func NewEmployeeRepository(connectionString string) *EmployeeRepository {
	return &EmployeeRepository{connectionString: connectionString}
}

func (e *EmployeeRepository) GetById(id int) *Models.Employee {
	return &Models.Employee{}
}

func (e *EmployeeRepository) GetAll() []*Models.Employee {
	return make([]*Models.Employee, 10)
}

func (e *EmployeeRepository) Create(employee *Models.Employee) *Models.Employee {
	return &Models.Employee{}
}

func (e *EmployeeRepository) Delete(id int) error {
	return *new(error)
}
