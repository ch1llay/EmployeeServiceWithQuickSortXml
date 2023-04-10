package handler

import (
	"EmployeeServiceWithQuickSortXml/Model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

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
func (h *Handler) GetAllEmployeesFull(w http.ResponseWriter, r *http.Request) {
	employees, err := h.EmployeeService.GetAllEmployeesFull()
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

func (h *Handler) GetEmployeeByIdWithBinarySearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if id, err := strconv.Atoi(vars["id"]); err == nil {
		if employee, err := h.EmployeeService.GetByIdWithBinarySearch(id); err == nil {
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
