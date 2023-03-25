package repository

import (
	"EmployeeServiceWithQuickSortXml/Models"
	"errors"
)

type EmployeeRepository struct {
	connectionString string
}

func NewEmployeeRepository(connectionString string) *EmployeeRepository {
	return &EmployeeRepository{connectionString: connectionString}
}

func (e *EmployeeRepository) GetById(id int) (*Models.Employee, error) {
	return &Models.Employee{}, errors.New("")
}

func (e *EmployeeRepository) GetAll() []*Models.Employee {
	return make([]*Models.Employee, 10)
}

func (e *EmployeeRepository) Create(employee *Models.Employee) (*Models.Employee, error) {
	return &Models.Employee{}, errors.New("")
}

func (e *EmployeeRepository) Update(employee *Models.Employee) (*Models.Employee, error) {
	return &Models.Employee{}, errors.New("")
}

func (e *EmployeeRepository) Delete(id int) error {
	return *new(error)
}
