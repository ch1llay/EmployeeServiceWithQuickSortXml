package repository

import (
	"EmployeeServiceWithQuickSortXml/Model"
	"EmployeeServiceWithQuickSortXml/internal/repository/query"
	"database/sql"
	"fmt"
)

func NewReportRepository(connectionString string) *ReportRepository {
	return &ReportRepository{
		ConnectionString: connectionString}
}
func (r *ReportRepository) Insert(report *Model.Report) (*Model.Report, error) {
	db, err := sql.Open("postgres", r.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	err = db.QueryRow(query.InsertReport, report.Name, report.Text, report.EmployeeId).Scan(&report.Id)
	if err != nil {
		return &Model.Report{}, err
	}

	return report, nil
}

func (r *ReportRepository) GetById(id int) (*Model.Report, error) {
	db, err := sql.Open("postgres", r.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	report := Model.Report{}
	err = db.QueryRow(query.GetByIdReport, id).Scan(&report.Id, &report.Name, &report.Text)
	if err != nil {
		return &Model.Report{}, err
	}

	return &report, nil
}

func (r *ReportRepository) GetByEmployeeId(employeeId int) (reports []*Model.Report, err error) {
	db, err := sql.Open("postgres", r.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	rows, err := db.Query(query.GetByEmployeeIdReport, employeeId)
	if err != nil {
		return
	}

	reports, err = r.scanReports(rows, reports)
	return
}
func (r *ReportRepository) DeleteById(id int) (deletingId int, err error) {
	db, err := sql.Open("postgres", r.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	err = db.QueryRow(query.DeleteByIdReport, id).Scan(&deletingId)
	return
}
func (r *ReportRepository) scanReports(rows *sql.Rows, reports []*Model.Report) ([]*Model.Report, error) {
	for rows.Next() {
		report := new(Model.Report)
		err := rows.Scan(&report.Id, &report.Name, &report.Text, &report.EmployeeId)
		if err != nil {
			fmt.Println(err)
			continue
		}

		reports = append(reports, report)
	}

	return reports, nil
}
