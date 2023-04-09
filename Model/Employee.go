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
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Text       string `json:"text"`
	EmployeeId int    `json:"employee_id"`
}

type EmployeeFull struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Lastname   string    `json:"lastname"`
	Patronymic string    `json:"patronymic"`
	Birthday   time.Time `json:"birthday"`
	Reports    []*Report `json:"reports"`
}
