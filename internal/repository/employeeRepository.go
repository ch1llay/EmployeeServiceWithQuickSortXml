package repository

import (
	"EmployeeServiceWithQuickSortXml/Model"
	"EmployeeServiceWithQuickSortXml/internal/repository/query"
	"EmployeeServiceWithQuickSortXml/migrations"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewEmployeeRepository(connectionString string) *EmployeeRepository {
	return &EmployeeRepository{
		ConnectionString: connectionString,
		//ReportRepository_: &ReportRepository{ConnectionString: connectionString},
	}
}
func InitRepository(connectionString string) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	var id int
	err = db.QueryRow(migrations.Initial).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}
}

func (e *EmployeeRepository) Insert(employee *Model.Employee) (*Model.Employee, error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	err = db.QueryRow(query.InsertEmployee, employee.Name, employee.Lastname, employee.Patronymic, employee.Birthday).Scan(&employee.Id)
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
	err = db.QueryRow(query.GetByIdEmployee, id).Scan(&employee.Id, &employee.Name, &employee.Lastname, &employee.Patronymic, &employee.Birthday)
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
	rows, err := db.Query(query.GetAllEmployees)
	defer rows.Close()

	if err != nil {
		return
	}

	employees, err = e.scanEmployees(rows, employees)

	return
}

func (e *EmployeeRepository) Update(employee *Model.Employee) (employeeRes *Model.Employee, err error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		return
	}

	defer db.Close()
	err = db.QueryRow(query.UpdateByIdEmployee, employee.Id, employee.Name, employee.Lastname, employee.Patronymic, employee.Birthday).Scan()
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

func (e *EmployeeRepository) scanEmployees(rows *sql.Rows, employees []*Model.Employee) ([]*Model.Employee, error) {
	for rows.Next() {
		employee := new(Model.Employee)
		err := rows.Scan(&employee.Id, &employee.Name, &employee.Lastname, &employee.Patronymic, &employee.Birthday)
		if err != nil {
			fmt.Println(err)
			continue
		}

		employees = append(employees, employee)

	}

	return employees, nil

}
