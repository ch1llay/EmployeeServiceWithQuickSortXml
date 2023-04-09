package repository

import "EmployeeServiceWithQuickSortXml/Model"

type PassportRep interface {
	Insert(employee *Model.Passport) (*Model.Passport, error)
	GetById(id int) (*Model.Passport, error)
	Get() ([]*Model.Passport, error)
	Update(newEmployee *Model.Employee) (*Model.Employee, error)
	DeleteById(id int) error
}

type PassportRepository struct {
	ConnectionString string
}

func (p PassportRepository) Insert(employee *Model.Passport) (*Model.Passport, error) {
	//TODO implement me
	panic("implement me")
}

func (p PassportRepository) GetById(id int) (*Model.Passport, error) {
	//TODO implement me
	panic("implement me")
}

func (p PassportRepository) Get() ([]*Model.Passport, error) {
	//TODO implement me
	panic("implement me")
}

func (p PassportRepository) Update(newEmployee *Model.Employee) (*Model.Employee, error) {
	//TODO implement me
	panic("implement me")
}

func (p PassportRepository) DeleteById(id int) error {
	//TODO implement me
	panic("implement me")
}
