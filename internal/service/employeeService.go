package service

import (
	"EmployeeServiceWithQuickSortXml/Model"
	"EmployeeServiceWithQuickSortXml/internal/repository"
	"errors"
)

type EmployeeServ interface {
	GetById(id int) (*Model.Employee, error)
	GetAll() (*Model.Employee, error)
	Create(employee Model.Employee) (*Model.Employee, error)
	DeleteById(id int) error
	GetAllSort() (*Model.Employee, error)
	UpdateById(employee *Model.Employee) (*Model.Employee, error)
}

type EmployeeService struct {
	EmployeeRepository repository.EmployeeRep
	FileRepository     repository.FileRep
}

func NewEmployeeService(employeeRep repository.EmployeeRep) *EmployeeService {
	return &EmployeeService{EmployeeRepository: employeeRep}
}

func (e *EmployeeService) GetById(id int) (*Model.Employee, error) {
	return e.EmployeeRepository.GetById(id)
}

func (e *EmployeeService) GetAll() ([]*Model.Employee, error) {
	return e.EmployeeRepository.Get()
}

func (e *EmployeeService) Create(employee *Model.Employee) (*Model.Employee, error) {
	return e.EmployeeRepository.Insert(employee)
}

func (e *EmployeeService) Update(employee *Model.Employee) (*Model.Employee, error) {
	_, err := e.GetById(employee.Id)
	if err != nil {
		return &Model.Employee{}, errors.New("404")
	}
	//TODO: пустые поля заменить на поля из old employee
	newEmployee, err := e.EmployeeRepository.Update(employee)
	return newEmployee, err
}

func (e *EmployeeService) Delete(id int) error {
	return e.EmployeeRepository.DeleteById(id)
}
