package handler

import (
	"EmployeeServiceWithQuickSortXml/Model"
	"EmployeeServiceWithQuickSortXml/internal/service"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	Router          *mux.Router
	EmployeeService service.EmployeeServ
	FileService     service.FileServ
}

func (h *Handler) InitRoutes() {
	h.Router.HandleFunc("ping", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "pong")
	})
	h.Router.HandleFunc("employee/{id}", h.GetEmployeeById).Methods("GET")
	h.Router.HandleFunc("employee/", h.GetAllEmployees).Methods("GET")
	h.Router.HandleFunc("employee/", h.CreateEmployee).Methods("POST")
	h.Router.HandleFunc("employee/", h.UpdateEmployee).Methods("PUT")
	h.Router.HandleFunc("employee/{id}", h.DeleteEmployee).Methods("DELETE")

	h.Router.HandleFunc("file/{guid}", h.GetFileById).Methods("GET")
	h.Router.HandleFunc("file/", h.GetAllFiles).Methods("GET")
}
func (h *Handler) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
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

func NewHandler(employeeService service.EmployeeServ, fileService service.FileServ) *Handler {
	h := &Handler{FileService: fileService, EmployeeService: employeeService, Router: mux.NewRouter()}
	h.InitRoutes()
	return h
}

func (h *Handler) GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if id, err := strconv.Atoi(vars["id"]); err == nil {
		if employee, err := h.EmployeeService.GetById(id); err == nil {
			h.respond(w, r, 200, employee)
		} else if err.Error() == "404" {
			h.respond(w, r, 404, "employee is not found")
		}
	} else {
		h.respond(w, r, 400, "BadRequest")
	}
}

func (h *Handler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees := h.EmployeeService.GetAll()
	h.respond(w, r, 200, employees)
}

func (h *Handler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee Model.Employee
	if err := h.GetJsonModel(r, &employee); err == nil {
		if newEmployee, errM := h.EmployeeService.Create(&employee); errM == nil {
			h.respond(w, r, 200, newEmployee.Id)
		} else {
			h.respond(w, r, 500, err)
		}
	} else {
		h.respond(w, r, 400, "")
	}

}

func (h *Handler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee Model.Employee
	if err := h.GetJsonModel(r, &employee); err == nil {
		newEmployee, errM := h.EmployeeService.Update(&employee)
		if errM.Error() == "404" {
			h.respond(w, r, 404, "employee is not found")
			return
		}
		h.respond(w, r, 200, newEmployee)
	} else {
		h.respond(w, r, 400, "")
	}
}

func (h *Handler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if id, err := strconv.Atoi(vars["id"]); err == nil {
		if errM := h.EmployeeService.Delete(id); errM != nil {
			h.respond(w, r, 200, "")
		} else if errM.Error() == "404" {
			h.respond(w, r, 404, "employee is not found")
		} else {
			h.respond(w, r, 500, "")
		}
	} else {
		h.respond(w, r, 400, "BadRequest")
	}
}

func (h *Handler) GetFileById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if len(strings.TrimSpace(id)) != 0 {
		if file, err := h.FileService.GetByGuid(id); err == nil {
			h.respond(w, r, 200, file) // todo: отправлять тип ответа File
		} else if err.Error() == "404" {
			h.respond(w, r, 404, "file is not found") // todo: отправлять тип ответа File
		} else {
			h.respond(w, r, 500, err) // todo: отправлять тип ответа File
		}
	} else {
		h.respond(w, r, 400, "BadRequest")
	}
}

func (h *Handler) GetAllFiles(w http.ResponseWriter, r *http.Request) {
	files := h.FileService.GetAll()
	h.respond(w, r, 200, files)
}
