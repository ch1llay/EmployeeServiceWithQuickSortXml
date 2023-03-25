package service

import (
	"EmployeeServiceWithQuickSortXml/Models"
	"EmployeeServiceWithQuickSortXml/internal/repository"
	"errors"
)

type Service interface {
	GetById(id int) (*Models.Employee, error)
	GetAll() *Models.Employee
	Create(employee Models.Employee) (*Models.Employee, error)
	Delete(id int) error
	GetAllSort() (*Models.Employee, error)
	Update(employee *Models.Employee) (*Models.Employee, error)
}

type EmployeeService struct {
	EmployeeRepository *repository.EmployeeRepository
	FileRepository     *repository.FileRepository
}

func NewEmployeeService(employeeRep *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{EmployeeRepository: employeeRep}
}

func (e *EmployeeService) GetById(id int) (*Models.Employee, error) {
	return e.EmployeeRepository.GetById(id)
}

func (e *EmployeeService) GetAll() []*Models.Employee {
	return e.EmployeeRepository.GetAll()
}

func (e *EmployeeService) Create(employee *Models.Employee) (*Models.Employee, error) {
	return e.EmployeeRepository.Create(employee)
}

func (e *EmployeeService) Update(employee *Models.Employee) (*Models.Employee, error) {
	_, err := e.GetById(employee.Id)
	if err != nil {
		return &Models.Employee{}, errors.New("404")
	}
	//TODO: пустые поля заменить на поля из old employee
	newEmployee, err := e.EmployeeRepository.Update(employee)
	return newEmployee, err
}

func (e *EmployeeService) Delete(id int) error {
	return e.EmployeeRepository.Delete(id)
}
