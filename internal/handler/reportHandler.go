package handler

import (
	"EmployeeServiceWithQuickSortXml/Model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *Handler) CreateReportForEmployee(w http.ResponseWriter, r *http.Request) {
	var report Model.Report
	if err := h.GetJsonModel(r, &report); err == nil {
		if newEmployee, errM := h.ReportService.CreateReportForEmployee(&report); errM == nil {
			h.responseModel(w, newEmployee)
			return
		} else {
			h.responseError(w, 500, err.Error())
			return
		}
	} else {
		h.responseError(w, 400, err.Error())
		return
	}
}

func (h *Handler) GetReportById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if id, err := strconv.Atoi(vars["id"]); err == nil {
		if report, err := h.ReportService.GetReportById(id); err == nil {
			h.responseModel(w, report)
			return
		} else if err.Error() == "404" {
			h.responseError(w, 404, "employee not founded")
			return
		} else {
			h.responseError(w, 500, err.Error())
			return
		}
	} else {
		h.responseError(w, 400, err.Error())
	}
	return
}
