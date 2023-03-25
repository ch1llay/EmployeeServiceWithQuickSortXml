package handler

import (
	"EmployeeServiceWithQuickSortXml/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	Router          *mux.Router
	EmployeeService *service.EmployeeService
	FileService     *service.FileService
}

func (h *Handler) InitRoutes() {
	h.Router.HandleFunc("/getById/{id}", h.GetEmployeeById)
}
func (h *Handler) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func NewHandler(employeeService *service.EmployeeService, fileService *service.FileService) *Handler {
	return &Handler{FileService: fileService, EmployeeService: employeeService}
}

func (h *Handler) GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if id, err := strconv.Atoi(vars["id"]); err == nil {
		h.respond(w, r, 200, h.EmployeeService.GetById(id))
	} else {
		h.respond(w, r, 400, "BadRequest")
	}
}
