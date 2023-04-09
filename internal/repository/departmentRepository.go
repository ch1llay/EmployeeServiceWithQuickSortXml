package repository

import "EmployeeServiceWithQuickSortXml/Model"

type Repository interface {
	Insert(employee *Model.) (*Model.Employee, error)
	GetById(id int) (*Model.Employee, error)
	Get() ([]*Model.Employee, error)
	Update(newEmployee *Model.Employee) (*Model.Employee, error)
	DeleteById(id int) error
}
