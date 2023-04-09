package Model

import "time"

type Employee struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Lastname   string    `json:"lastname"`
	Patronymic string    `json:"patronymic"`
	Birthday   time.Time `json:"birthday"`
	Passport *Passport `json:"passport"`
	Departments []int `json:"departments"`
}

type Passport struct {
	Id         int    `json:"id"`
	Number     string `json:"number"`
	EmployeeId int    `json:"employee_id"`
}

type Department
