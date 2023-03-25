package repository

import (
	"EmployeeServiceWithQuickSortXml/Models"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
)

type EmployeeRepository struct {
	connectionString string
	db               *sql.DB
}

func (e *EmployeeRepository) Open() error {
	db, err := sql.Open("postgres", e.connectionString)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	e.db = db

	e.db.Exec()

	return nil
}

func (e *EmployeeRepository) Close() {
	e.db.Close()
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
	err := e.db.QueryRow("insert into employees (name, age, department_name) values ($1, $2, $3) RETURNING id", employee.Name, employee.Age, employee.DepartmentName).Scan(&employee.Id)
	if err != nil {
		return &Models.Employee{}, err
	}

	return employee, nil

}

func (e *EmployeeRepository) Update(employee *Models.Employee) (*Models.Employee, error) {

	err := e.db.QueryRow("update employees set name = $2 set age = $3 set department_name = $4 where id = $1", employee.Name, employee.Age, employee.DepartmentName)
	if err != nil {
		return employee, err
	}

	return employee, nil

}

func (e *EmployeeRepository) Delete(id int) error {
	var id int
	err := e.db.QueryRow("insert into employees (name, age, department_name) values ($1, $2, $3) RETURNING id", employee.Name, employee.Age, employee.DepartmentName).Scan(&id)
	if err != nil {
		return &Models.Employee{}, err
	}

	employee.Id = id
	return employee, nil
}
