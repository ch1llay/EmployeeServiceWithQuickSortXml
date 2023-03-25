package Models

type Employee struct {
	Id             int       `json: "id"`
	Name           string    `json:"name"`
	Age            int       `json:"age"`
	Passport       *Passport `json:"passport"`
	DepartmentName string    `json:"departmentName"`
}

type Passport struct {
	Number string
}
