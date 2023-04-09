package Model

import "time"

type Employee struct {
	Id         int       `json:"id" xml:"id,attr"`
	Name       string    `json:"name" xml:"name,attr"`
	Lastname   string    `json:"lastname" xml:"lastname,attr"`
	Patronymic string    `json:"patronymic" xml:"patronymic,attr"`
	Birthday   time.Time `json:"birthday" xml:"birthday"`
}

type Passport struct {
	Id         int    `json:"id"`
	Number     string `json:"number"`
	EmployeeId int    `json:"employee_id"`
}

type PassportUpdateModel struct {
	Number string `json:"number"`
	Type   string `json:"type"`
}

type Department struct {
	Id          int `json:"id"`
	Name        string
	Description string `json:"description"`
}

type EmployeeFull struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Lastname    string       `json:"lastname"`
	Patronymic  string       `json:"patronymic"`
	Birthday    time.Time    `json:"birthday"`
	Passport    *Passport    `json:"passport"`
	Departments []Department `json:"departments"`
}

type EmployeeUpdateModel struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	Lastname       string    `json:"lastname"`
	Patronymic     string    `json:"patronymic"`
	Birthday       time.Time `json:"birthday"`
	PassportNumber string    `json:"passport_number"`
	Departments    []int     `json:"departments"`
}
