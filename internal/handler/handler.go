package handler

import (
	"EmployeeServiceWithQuickSortXml/Model"
	"EmployeeServiceWithQuickSortXml/internal/service"
	"EmployeeServiceWithQuickSortXml/pkg/XMLHelper"
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
	h.Router.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "pong")
	})
	h.Router.HandleFunc("/employee/{id}", h.GetEmployeeById).Methods("GET")
	h.Router.HandleFunc("/employee/", h.GetAllEmployees).Methods("GET")
	h.Router.HandleFunc("/employee/", h.CreateEmployee).Methods("POST")
	h.Router.HandleFunc("/employee/", h.UpdateEmployee).Methods("PUT")
	h.Router.HandleFunc("/employee/{id}", h.DeleteEmployee).Methods("DELETE")

	h.Router.HandleFunc("/file/get", h.GetXMLFileId).Methods("GET")
	h.Router.HandleFunc("/file/get-sorting", h.GetXMLFileIdSorting).Methods("GET")
	h.Router.HandleFunc("/file/{guid}", h.GetFileById).Methods("GET")
}
func (h *Handler) respond(w http.ResponseWriter, code int, response interface{}) {
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(code)
	}

	w.WriteHeader(500)
}

func (h *Handler) responseModel(w http.ResponseWriter, response interface{}) {
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

func NewHandler(employeeService service.EmployeeServ, fileService service.FileServ) *Handler {
	h := &Handler{FileService: fileService, EmployeeService: employeeService, Router: mux.NewRouter()}
	h.InitRoutes()
	return h
}

func (h *Handler) GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if id, err := strconv.Atoi(vars["id"]); err == nil {
		if employee, err := h.EmployeeService.GetById(id); err == nil {
			h.responseModel(w, employee)
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

func (h *Handler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := h.EmployeeService.GetAll()
	if err != nil {
		h.responseError(w, 500, err.Error())
		return
	}
	h.responseModel(w, employees)
	return
}

func (h *Handler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee Model.Employee
	if err := h.GetJsonModel(r, &employee); err == nil {
		if newEmployee, errM := h.EmployeeService.Create(&employee); errM == nil {
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

func (h *Handler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee Model.Employee
	if err := h.GetJsonModel(r, &employee); err == nil {
		newEmployee, errM := h.EmployeeService.Update(&employee)
		if errM.Error() == "404" {
			h.responseError(w, 404, "not found")
			return
		}
		h.responseModel(w, newEmployee)
		return
	} else {
		w.WriteHeader(400)
		return
	}
}

func (h *Handler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if id, err := strconv.Atoi(vars["id"]); err == nil {
		if delettingId, errM := h.EmployeeService.Delete(id); errM != nil && delettingId == id {
			w.WriteHeader(200)
		} else if errM.Error() == "404" {
			w.WriteHeader(404)
			return
		} else {
			h.responseError(w, 500, err.Error())
			return
		}
	} else {
		w.WriteHeader(400)
		return
	}
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

func (h *Handler) GetXMLFileId(w http.ResponseWriter, r *http.Request) {
	employees, err := h.EmployeeService.GetAll()
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

	h.responseModel(w, &Model.FileResponse{Guid: fileGuid})

}

func (h *Handler) GetXMLFileIdSorting(w http.ResponseWriter, r *http.Request) {
	employees, err := h.EmployeeService.GetAllSortByBirthday()
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

	h.responseModel(w, Model.FileResponse{Guid: fileGuid})
}
