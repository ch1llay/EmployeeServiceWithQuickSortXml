package repository

import (
	"EmployeeServiceWithQuickSortXml/Model"
	"EmployeeServiceWithQuickSortXml/internal/repository/query"
	"database/sql"
	"fmt"
)

type DepartmentRep interface {
	Insert(department *Model.Department) (*Model.Department, error)
	GetById(id int) (*Model.Department, error)
	Get() ([]*Model.Department, error)
	Update(newDepartment *Model.Department) (*Model.Department, error)
	DeleteById(id int) (int, error)
}

type DepartmentRepository struct {
	ConnectionString string
}

func (e *DepartmentRepository) Insert(department *Model.Department) (*Model.Department, error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	err = db.QueryRow(query.InsertDepartment).Scan(&department.Id)
	if err != nil {
		return &Model.Department{}, err
	}

	return department, nil

}
func (e *DepartmentRepository) GetById(id int) (*Model.Department, error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	department := Model.Department{}
	err = db.QueryRow(query.GetByIdDepartment, department.Id).Scan(&department.Name, &department.Description)
	if err != nil {
		return &Model.Department{}, err
	}

	return &department, nil
}

func (e *DepartmentRepository) Get() (departments []*Model.Department, err error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	rows, err := db.Query(query.GetByIdDepartment)
	defer rows.Close()

	if err != nil {
		return
	}

	e.scanDepartments(rows, departments)

	return
}

func (e *DepartmentRepository) Update(department *Model.Department) (departmentRes *Model.Department, err error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		return
	}

	defer db.Close()
	err = db.QueryRow(query.UpdateByIdDepartment, department.Id, department.Name).Scan()
	if err != nil {
		departmentRes = department
	}

	return

}

func (e *DepartmentRepository) DeleteById(id int) (deletingId int, err error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	err = db.QueryRow(query.DeleteByIdDepartment, id).Scan(&deletingId)
	return
}

func (e *DepartmentRepository) scanDepartments(rows *sql.Rows, departments []*Model.Department) {
	for rows.Next() {
		department := new(Model.Department)
		err := rows.Scan(&department.Id, &department.Name, &department.Description)
		if err != nil {
			fmt.Println(err)
		}
		departments = append(departments, department)
	}
}
