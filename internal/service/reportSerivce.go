package service

import (
	"EmployeeServiceWithQuickSortXml/Model"
	"EmployeeServiceWithQuickSortXml/internal/repository"
)

type ReportServ interface {
	CreateReportForEmployee(report *Model.Report) (*Model.Report, error)
	GetReportById(id int) (*Model.Report, error)
	GetReportsByEmployeeId(id int) ([]*Model.Report, error)
	DeleteById(id int) (int, error)
}

type ReportService struct {
	ReportRepository repository.ReportRep
}

func (r *ReportService) CreateReportForEmployee(report *Model.Report) (*Model.Report, error) {
	return r.ReportRepository.Insert(report)
}

func (r *ReportService) GetReportById(id int) (*Model.Report, error) {
	return r.ReportRepository.GetById(id)
}

func (r *ReportService) GetReportsByEmployeeId(id int) ([]*Model.Report, error) {
	return r.ReportRepository.GetByEmployeeId(id)
}

func (e *ReportService) DeleteById(id int) (int, error) {
	return e.ReportRepository.DeleteById(id)
}
