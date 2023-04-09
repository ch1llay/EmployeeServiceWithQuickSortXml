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
}
