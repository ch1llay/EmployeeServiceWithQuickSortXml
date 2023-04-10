package repository

import (
	"EmployeeServiceWithQuickSortXml/Model"
)

type EmployeeRep interface {
	Insert(employee *Model.Employee) (*Model.Employee, error)
	GetById(id int) (*Model.Employee, error)
	Get() ([]*Model.Employee, error)
	Update(newEmployee *Model.Employee) (*Model.Employee, error)
	DeleteById(id int) (int, error)
}
type EmployeeRepository struct {
	ConnectionString string
	//ReportRepository_ ReportRep
	//SqlFileReader        *SqlReader
}

type ReportRep interface {
	Insert(report *Model.Report) (*Model.Report, error)
	GetById(id int) (*Model.Report, error)
	GetByEmployeeId(employeeId int) ([]*Model.Report, error)
	DeleteById(id int) (int, error)
}

type ReportRepository struct {
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
