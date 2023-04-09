package repository

import (
	"EmployeeServiceWithQuickSortXml/Model"
	"EmployeeServiceWithQuickSortXml/internal/repository/query"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type EmployeeRep interface {
	Insert(employee *Model.Employee) (*Model.Employee, error)
	GetById(id int) (*Model.Employee, error)
	Get() ([]*Model.Employee, error)
	Update(newEmployee *Model.Employee) (*Model.Employee, error)
	DeleteById(id int) error
}
type EmployeeRepository struct {
	ConnectionString   string
	passportRepository EmployeeRep
	departmentRepository Repository
	SqlFileReader    *SqlReader
}

func NewEmployeeRepository(connectionString string, scriptPath string) *EmployeeRepository {
	return &EmployeeRepository{ConnectionString: connectionString}}
}

func (e *EmployeeRepository) Insert(employee *Model.Employee) (*Model.Employee, error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	err = db.QueryRow(query.InsertEmployee).Scan(&employee.Id)
	if err != nil {
		return &Model.Employee{}, err
	}

	return employee, nil

}
func (e *EmployeeRepository) GetById(id int) (*Model.Employee, error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	employee := Model.Employee{}
	err = db.QueryRow(query.GetByIdEmployee, employee.Id).Scan(&employee.Name, &employee.Lastname, &employee.Patronymic, &employee.Birthday)
	if err != nil {
		return &Model.Employee{}, err
	}

	return &employee, nil
}

func (e *EmployeeRepository) Get() (employees []*Model.Employee, err error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	rows, err := db.Query(query.GetByIdEmployee)
	defer rows.Close()


	if err != nil {
		return
	}

	e.scanEmployees(rows, employees)
}

func (e *EmployeeRepository) Update(employee *Model.Employee) (employeeRes *Model.Employee, err error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		return
	}

	defer db.Close()
	err = db.QueryRow(query.UpdateByIdEmployee, employee.Id, employee.Name, employee.Lastname, employee.Patronymic).Scan()
	if err != nil {
		employeeRes = employee
	}

	return

}

func (e *EmployeeRepository) DeleteById(id int) (deletingId int, err error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	err = db.QueryRow(query.DeleteByIdEmployee, id).Scan(&deletingId)
	return
}

func (e *EmployeeRepository) scanEmployees(rows *sql.Rows, employees []*Model.Employee) []Model.Employee {
	for rows.Next() {
		employee := Model.Employee{}
		err := rows.Scan(&employee.Id, &employee.Name, &employee.Lastname, &employee.Patronymic, &employee.Birthday)
		if err != nil {
			fmt.Println(err)
			continue
		}

		employees = append(employees, employee)
	}

	return employees
}
