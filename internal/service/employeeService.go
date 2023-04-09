package service

import (
	"EmployeeServiceWithQuickSortXml/Model"
	"EmployeeServiceWithQuickSortXml/internal/repository"
	"errors"
	"fmt"
)

type EmployeeServ interface {
	GetById(id int) (*Model.Employee, error)
	GetAll() ([]*Model.Employee, error)
	GetEmployeeFullById(id int) (*Model.EmployeeFull, error)
	GetAllEmployeesFull() ([]*Model.EmployeeFull, error)
	Create(employee *Model.Employee) (*Model.Employee, error)
	Delete(id int) (int, error)
	GetAllSort() ([]*Model.Employee, error)
	GetAllEmployeesFullSort() ([]*Model.EmployeeFull, error)
	Update(employee *Model.Employee) (*Model.Employee, error)
}

type EmployeeService struct {
	EmployeeRepository repository.EmployeeRep
	ReportRepository   repository.ReportRep
	FileRepository     repository.FileRep
}

func NewEmployeeService(employeeRep repository.EmployeeRep) *EmployeeService {
	return &EmployeeService{EmployeeRepository: employeeRep}
}

func (e *EmployeeService) GetById(id int) (*Model.Employee, error) {
	return e.EmployeeRepository.GetById(id)
}

func (e *EmployeeService) GetEmployeeFullById(id int) (*Model.EmployeeFull, error) {
	employee, err := e.EmployeeRepository.GetById(id)
	if err != nil {
		return &Model.EmployeeFull{}, errors.New("404")
	}
	reports, err := e.ReportRepository.GetByEmployeeId(id)
	if err != nil {
		return &Model.EmployeeFull{}, err
	}
	employeeFull := employee.ToEmployeeFull(reports)
	return employeeFull, nil
}

func (e *EmployeeService) GetAll() ([]*Model.Employee, error) {
	return e.EmployeeRepository.Get()
}

func (e *EmployeeService) GetAllEmployeesFull() ([]*Model.EmployeeFull, error) {
	employees, err := e.EmployeeRepository.Get()
	if err != nil {
		return nil, err
	}
	employessFull := make([]*Model.EmployeeFull, len(employees))
	for _, employee := range employees {
		reports, err := e.ReportRepository.GetByEmployeeId(employee.Id)
		if err != nil {
			fmt.Println(err)
			continue
		}

		employessFull = append(employessFull, employee.ToEmployeeFull(reports))
	}

	return employessFull, nil
}

func (e *EmployeeService) GetAllSort() ([]*Model.Employee, error) {
	return e.EmployeeRepository.Get()
	// todo:quicksort(employeess)
}

func (e *EmployeeService) GetAllEmployeesFullSort() ([]*Model.EmployeeFull, error) {
	return e.GetAllEmployeesFull()
	// todo:quicksort(employeess)
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

func (e *EmployeeService) Delete(id int) (int, error) {
	return e.EmployeeRepository.DeleteById(id)
}
