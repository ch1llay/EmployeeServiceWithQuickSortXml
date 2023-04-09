package XMLHelper

import (
	"EmployeeServiceWithQuickSortXml/Model"
	"encoding/xml"
	"time"
)

func GetXmlFile(employees []*Model.Employee) (*Model.File, error) {
	tmp := struct {
		Employees []*Model.Employee `xml:"employee"`
		XMLName   xml.Name          `xml:"Employees"`
	}{Employees: employees}
	data, err := xml.MarshalIndent(tmp, "", "   ")
	return &Model.File{FileName: "file.xml", Data: data, InsertDate: time.Now()}, err
}
