package Model

func (e *Employee) ToEmployeeFull(reports []*Report) *EmployeeFull {
	return &EmployeeFull{
		Id:         e.Id,
		Name:       e.Name,
		Lastname:   e.Lastname,
		Patronymic: e.Patronymic,
		Birthday:   e.Birthday,
		Reports:    reports,
	}
}
