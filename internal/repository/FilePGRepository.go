package repository

import (
	"EmployeeServiceWithQuickSortXml/Model"
	"EmployeeServiceWithQuickSortXml/internal/repository/query"
	"database/sql"
	"errors"
	"fmt"
)

func NewFilePGRepository(connectionString string) *FilePGRepository {
	return &FilePGRepository{ConnectionString: connectionString}
}
func (e *FilePGRepository) Insert(file *Model.File) (string, error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	err = db.QueryRow(query.InsertFile, file.FileName, file.InsertDate, file.Data).Scan(&file.Id)
	if err != nil {
		return "", err
	}

	return file.Id, nil

}
func (e *FilePGRepository) GetById(id string) (*Model.File, error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
		return &Model.File{}, err
	}

	defer db.Close()
	file := Model.File{}
	err = db.QueryRow(query.GetByIdFile, id).Scan(&file.Id, &file.FileName, &file.InsertDate, &file.Data)
	if err != nil {
		return &Model.File{}, errors.New("404")
	}

	return &file, nil
}

func (e *FilePGRepository) DeleteById(id string) (err error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	err = db.QueryRow(query.DeleteByIdFile, id).Scan()
	return
}
