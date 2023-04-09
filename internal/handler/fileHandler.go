package handler

import (
	"EmployeeServiceWithQuickSortXml/Model"
	"EmployeeServiceWithQuickSortXml/pkg/XMLHelper"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) GetXMLFileId(w http.ResponseWriter, r *http.Request) {
	employees, err := h.EmployeeService.GetAllEmployeesFull()
	if err != nil {
		w.WriteHeader(500)
	}
	file, err := XMLHelper.GetXmlFile(employees)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	fileGuid, err := h.FileService.Insert(file)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	h.responseModel(w, &Model.FileResponse{Guid: fileGuid, TypeSorting: "WithoutSorting"})

}

func (h *Handler) GetXMLFileIdSortingEmployeeFullByBirthday(w http.ResponseWriter, r *http.Request) {
	employees, err := h.EmployeeService.GetAllSortEmployeesFullByBirthday()
	if err != nil {
		w.WriteHeader(500)
	}
	file, err := XMLHelper.GetXmlFile(employees)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	fileGuid, err := h.FileService.Insert(file)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	h.responseModel(w, Model.FileResponse{Guid: fileGuid, TypeSorting: "ByBirthday"})
}

func (h *Handler) GetXMLFileIdSortingEmployeeFullByReportCount(w http.ResponseWriter, r *http.Request) {
	employees, err := h.EmployeeService.GetAllEmployeesFullSortByReportCount()
	if err != nil {
		w.WriteHeader(500)
	}
	file, err := XMLHelper.GetXmlFile(employees)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	fileGuid, err := h.FileService.Insert(file)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	h.responseModel(w, Model.FileResponse{Guid: fileGuid, TypeSorting: "ByReportCount"})
}

func (h *Handler) GetFileById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["guid"]
	if len(strings.TrimSpace(id)) != 0 {
		if file, err := h.FileService.GetById(id); err == nil {
			contentType := "application/xml"
			//contentType = "application/octet-stream"
			w.Header().Add("Content-Type", contentType)
			w.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, file.FileName))
			w.Header().Add("Content-Length", strconv.Itoa(len(file.Data)))
			w.Write(file.Data)
			return
		} else if err.Error() == "404" {
			w.WriteHeader(404)
		} else {
			w.WriteHeader(500)
		}
	} else {
		w.WriteHeader(400)
	}
}
