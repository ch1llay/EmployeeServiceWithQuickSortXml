package repository

import (
	"EmployeeServiceWithQuickSortXml/Model"
	"EmployeeServiceWithQuickSortXml/internal/repository/query"
	"database/sql"
	"fmt"
)

func (e *FilePGRepository) Insert(file *Model.File) (*Model.File, error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	err = db.QueryRow(query.InsertFile).Scan(&file.Id)
	if err != nil {
		return &Model.File{}, err
	}

	return file, nil

}
func (e *FilePGRepository) GetById(id int) (*Model.File, error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	file := Model.File{}
	err = db.QueryRow(query.GetByIdFile, file.Id).Scan(&file.FileName, &file.Data)
	if err != nil {
		return &Model.File{}, err
	}

	return &file, nil
}

func (e *FilePGRepository) DeleteById(id int) (deletingId int, err error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	err = db.QueryRow(query.DeleteByIdFile, id).Scan(&deletingId)
	return
}

func (e *FilePGRepository) scanFilePGs(rows *sql.Rows, files []*Model.File) {
	for rows.Next() {
		file := new(Model.File)
		err := rows.Scan(&file.Id, &file.FileName, &file.InsertDate)
		if err != nil {
			fmt.Println(err)
		}
		files = append(files, file)
	}
}
