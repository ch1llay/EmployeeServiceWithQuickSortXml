package repository

import "EmployeeServiceWithQuickSortXml/Model"

type EmployeeRep interface {
	Insert(employee *Model.Employee) (*Model.Employee, error)
	GetById(id int) (*Model.Employee, error)
	Get() ([]*Model.Employee, error)
	Update(newEmployee *Model.Employee) (*Model.Employee, error)
	DeleteById(id int) (int, error)
}
type EmployeeRepository struct {
	ConnectionString     string
	passportRepository   PassportRep
	departmentRepository DepartmentRep
	//SqlFileReader        *SqlReader
}

type PassportRep interface {
	Insert(employee *Model.Passport) (*Model.Passport, error)
	GetById(id int) (*Model.Passport, error)
	Get() ([]*Model.Passport, error)
	Update(newEmployee *Model.Employee) (*Model.Employee, error)
	DeleteById(id int) error
}

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

type PassportRepository struct {
	ConnectionString string
}
type FilePGRepository struct {
	ConnectionString string
}
type FileRep interface {
	Insert(file *Model.File) (string, error)
	GetById(guid string) (*Model.File, error)
	DeleteById(guid string) error
}

type FileMongoRepository struct {
	connectionString, databaseName, collectionName string
}
