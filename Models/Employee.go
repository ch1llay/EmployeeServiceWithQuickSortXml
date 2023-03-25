package Models

type Employee struct {
	Name           string    `json: "-"`
	Passport       *Passport `json:"passport"`
	DepartmentName string    `json:"departmentName"`
}

type Passport struct {
	Number string
}
