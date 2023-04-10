package handler

import (
	"EmployeeServiceWithQuickSortXml/internal/service"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	Router          *mux.Router
	EmployeeService service.EmployeeServ
	ReportService   service.ReportServ
	FileService     service.FileServ
}

func (h *Handler) InitRoutes() {
	h.Router.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "pong")
	})
	h.Router.HandleFunc("/employee/{id:int}", h.GetEmployeeById).Methods("GET")
	h.Router.HandleFunc("/employee/", h.GetAllEmployees).Methods("GET")
	h.Router.HandleFunc("/employee/full", h.GetAllEmployeesFull).Methods("GET")
	h.Router.HandleFunc("/employee/", h.CreateEmployee).Methods("POST")
	h.Router.HandleFunc("/employee/", h.UpdateEmployee).Methods("PUT")
	h.Router.HandleFunc("/employee/{id}", h.DeleteEmployee).Methods("DELETE")
	h.Router.HandleFunc("/employee/binary-search/{id}", h.GetEmployeeByIdWithBinarySearch).Methods("GET")

	h.Router.HandleFunc("/report/", h.CreateReportForEmployee).Methods("POST")
	h.Router.HandleFunc("/report/{id}", h.GetReportById).Methods("Get")

	h.Router.HandleFunc("/file/get", h.GetXMLFileId).Methods("GET")
	h.Router.HandleFunc("/file/generate/get-sorting-birthday", h.GetXMLFileIdSortingEmployeeFullByBirthday).Methods("GET")
	h.Router.HandleFunc("/file/generate/get-sorting-report-count", h.GetXMLFileIdSortingEmployeeFullByReportCount).Methods("GET")
	h.Router.HandleFunc("/file/{guid}", h.GetFileById).Methods("GET")
}
func (h *Handler) respond(w http.ResponseWriter, code int, response interface{}) {
	w.Header().Set("Content-Type", "application/json;encoding=utf-8")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(code)
	}

	w.WriteHeader(500)
}

func (h *Handler) responseModel(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json;encoding=utf-8")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(500)
	}

	w.WriteHeader(200)
}

func (h *Handler) responseError(w http.ResponseWriter, code int, text string) {
	w.Header().Set("Content-Type", "application/json;encoding=utf-8")
	err := json.NewEncoder(w).Encode(map[string]string{"error": text})
	if err != nil {
		w.WriteHeader(200)
	}

	w.WriteHeader(500)
}

func (h *Handler) GetJsonModel(r *http.Request, model interface{}) error {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(r.Body)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.NewDecoder(bytes.NewBuffer(b)).Decode(&model)
	return err
}

func NewHandler(employeeService service.EmployeeServ, reportService service.ReportServ, fileService service.FileServ) *Handler {
	h := &Handler{FileService: fileService, EmployeeService: employeeService, ReportService: reportService, Router: mux.NewRouter()}
	h.InitRoutes()
	return h
}
