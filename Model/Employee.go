package Model

import "time"

type Employee struct {
	Id         int       `json:"id" xml:"id,attr"`
	Name       string    `json:"name" xml:"name,attr"`
	Lastname   string    `json:"lastname" xml:"lastname,attr"`
	Patronymic string    `json:"patronymic" xml:"patronymic,attr"`
	Birthday   time.Time `json:"birthday" xml:"birthday"`
}

type Report struct {
	Id         int    `json:"id" xml:"id,attr"`
	Name       string `json:"name" xml:"name,attr"`
	Text       string `json:"text" xml:"text"`
	EmployeeId int    `json:"employee_id" xml:"employeeId,attr"`
}

type EmployeeFull struct {
	Id         int       `json:"id" xml:"id,attr"`
	Name       string    `json:"name" xml:"name,attr"`
	Lastname   string    `json:"lastname" xml:"lastname,attr"`
	Patronymic string    `json:"patronymic" xml:"patronymic,attr"`
	Birthday   time.Time `json:"birthday" xml:"birthday"`
	Reports    []*Report `json:"reports" xml:"report"`
}
