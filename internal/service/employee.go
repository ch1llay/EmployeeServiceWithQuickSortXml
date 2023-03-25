package service

import (
	"EmployeeServiceWithQuickSortXml/Models"
	"EmployeeServiceWithQuickSortXml/internal/repository"
)

type Service interface {
	GetById(id int) *Models.Employee
	GetAll() *Models.Employee
	Create(employee Models.Employee) *Models.Employee
	Delete(id int) bool
	GetAllSort() *Models.Employee
	Update(employee *Models.Employee) *Models.Employee
}

type EmployeeService struct {
	EmployeeRepository *repository.EmployeeRepository
	FileRepository     *repository.FileRepository
}

func NewEmployeeService(employeeRep *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{EmployeeRepository: employeeRep}
}

func (e *EmployeeService) GetById(id int) *Models.Employee {
	return e.EmployeeRepository.GetById(id)
}

func (e *EmployeeService) GetAll(id int) []*Models.Employee {
	return e.EmployeeRepository.GetAll()
}

func (e *EmployeeService) Create(employee *Models.Employee) *Models.Employee {
	return e.EmployeeRepository.Create(employee)
}

func (e *EmployeeService) Delete(id int) error {
	return e.EmployeeRepository.Delete(id)
}
