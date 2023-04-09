package XMLHelper

import (
	"EmployeeServiceWithQuickSortXml/Model"
	"encoding/xml"
	"time"
)

func GetXmlFile(employees []*Model.Employee) (*Model.File, error) {
	data, err := xml.Marshal(employees)
	return &Model.File{FileName: "filename1", Data: data, InsertDate: time.Now()}, err
}
